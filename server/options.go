package server

import "net/http"

type Middleware = func(h http.Handler) http.Handler

type EktoServer struct {
	gatewayAddr     string
	middlewareStack []Middleware
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
