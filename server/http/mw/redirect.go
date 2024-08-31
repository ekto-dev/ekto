package mw

import (
	"context"
	"net/http"

	"github.com/ekto-dev/ekto/protoc-gen-ekto/ekto"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/protobuf/proto"
)

func Redirect(ctx context.Context, response http.ResponseWriter, msg proto.Message) error {
	// Check for either a redirect gRPC header, or for the response to be an ekto.Redirect message
	// If the response is an ekto.Redirect message, set the Location header to the URL in the message
	// If the response is a redirect gRPC header, set the Location header to the URL in the header
	// If the response is neither, return nil
	if redirect, ok := msg.(*ekto.Redirect); ok {
		response.Header().Set("Location", redirect.Url)
		redirectCode := redirect.GetCode()
		if redirect.GetCode() == ekto.RedirectCode_Default {
			redirectCode = ekto.RedirectCode_TEMPORARY_REDIRECT
		}
		response.WriteHeader(int(redirectCode))

		return nil
	}

	md, ok := runtime.ServerMetadataFromContext(ctx)
	if !ok {
		return nil
	}

	if vals := md.HeaderMD.Get("Location"); len(vals) > 0 {
		response.Header().Set("Location", vals[0])
		response.WriteHeader(http.StatusTemporaryRedirect)
	}

	return nil
}
