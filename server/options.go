package server

type EktoServer struct {
	gatewayAddr string
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

func WithGateway(addr string) Option {
	return func(s *EktoServer) {
		s.gatewayAddr = addr
	}
}
