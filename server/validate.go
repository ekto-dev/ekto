package server

import (
	"buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	"context"
	"errors"
	"github.com/bufbuild/protovalidate-go"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"net/http"
)

const ErrCodeValidation = "validation_error"

// ValidateUnaryServerInterceptor returns a new unary server interceptor that validates incoming messages.
func ValidateUnaryServerInterceptor(validator *protovalidate.Validator) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp interface{}, err error) {
		switch msg := req.(type) {
		case proto.Message:
			if err = validator.Validate(msg); err != nil {
				s := status.New(codes.InvalidArgument, ErrCodeValidation)
				var valErr *protovalidate.ValidationError
				if errors.As(err, &valErr) {
					s, err := s.WithDetails(valErr.ToProto())
					if err == nil {
						return nil, s.Err()
					}
				}

				return nil, s.Err()
			}
		default:
			return nil, errors.New("unsupported message type")
		}

		return handler(ctx, req)
	}
}

func HTTPErrorHandler(ctx context.Context, mux *runtime.ServeMux, marshaler runtime.Marshaler, w http.ResponseWriter, req *http.Request, err error) {
	s := status.Convert(err)
	pb := s.Proto()
	if pb.Message != ErrCodeValidation {
		runtime.DefaultHTTPErrorHandler(ctx, mux, marshaler, w, req, err)
	}

	w.Header().Del("Trailer")
	w.Header().Del("Transfer-Encoding")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(422)

	if len(pb.Details) == 0 {
		// validation error with no details, so we return an empty 422
		w.Write(nil)
		return
	}

	violations := &validate.Violations{}
	if err := pb.Details[0].UnmarshalTo(violations); err != nil {
		return
	}

	b, err := marshaler.Marshal(violations)
	if err != nil {
		return
	}

	w.Write(b)
}
