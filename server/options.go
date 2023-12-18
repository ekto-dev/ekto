package server

import (
	"net/http"

	"github.com/bufbuild/protovalidate-go"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type Middleware = func(h http.Handler) http.Handler

type EktoServer struct {
	gatewayAddr     string
	middlewareStack []Middleware
	logger          *zap.Logger
}

type Option func(*EktoServer)

func NewEktoServer(opts ...Option) *EktoServer {
	s := &EktoServer{}
	for _, opt := range opts {
		opt(s)
	}
	return s
}

func (s *EktoServer) HasGateway() bool {
	return s.gatewayAddr != ""
}

func (s *EktoServer) GatewayAddr() string {
	return s.gatewayAddr
}

func (s *EktoServer) ApplyMiddleware(h http.Handler) http.Handler {
	for _, mw := range s.middlewareStack {
		h = mw(h)
	}

	return h
}

// Interceptors returns a set of default interceptors to apply
func (s *EktoServer) Interceptors() ([]grpc.UnaryServerInterceptor, error) {
	validator, err := protovalidate.New()
	if err != nil {
		return nil, err
	}

	mw := []grpc.UnaryServerInterceptor{
		otelgrpc.UnaryServerInterceptor(),
	}

	if s.logger != nil {
		mw = append(mw, LoggingInterceptor(s.logger))
	}

	mw = append(mw, ValidateUnaryServerInterceptor(validator))
	mw = append(mw, GormErrorInterceptor())
	return mw, nil
}

func WithGateway(addr string) Option {
	return func(s *EktoServer) {
		s.gatewayAddr = addr
	}
}

func WithMiddleware(mw ...Middleware) Option {
	return func(s *EktoServer) {
		s.middlewareStack = append(s.middlewareStack, mw...)
	}
}

func WithLogger(logger *zap.Logger) Option {
	return func(s *EktoServer) {
		s.logger = logger
	}
}
