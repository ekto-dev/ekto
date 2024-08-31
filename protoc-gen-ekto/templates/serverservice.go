package templates

const ServerServiceTpl = `
func Start{{ .Name }}Server(ctx context.Context, rpcListenAddr string, srv {{ name . }}, opts ...ektoserver.Option) error {
	ektoServer := ektoserver.NewEktoServer(opts...)
	interceptors, err := ektoServer.Interceptors()
	if err != nil {
		return err
	}

	server := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			interceptors...,
		),
	)
	Register{{ name . }}(server, srv)
	reflection.Register(server)

	eg, _ := errgroup.WithContext(ctx)
	eg.Go(func() error {
		lis, err := net.Listen("tcp", rpcListenAddr)

		if err != nil {
			return err
		}

		return server.Serve(lis)
	})

	{{- if hasGateway . }}
	if ektoServer.HasGateway() {
		eg.Go(func() error {
			conn, err := grpc.DialContext(
				ctx,
				rpcListenAddr,
				grpc.WithBlock(),
				grpc.WithTransportCredentials(insecure.NewCredentials()),
			)

			if err != nil {
				return err
			}

			gwmux := runtime.NewServeMux(
				ektoServer.GatewayMuxOpts()...,
			)

			err = Register{{ .Name }}Handler(ctx, gwmux, conn)
			if err != nil {
				return err
			}

			gwServer := &http.Server{
				Addr:    ektoServer.GatewayAddr(),
				Handler: ektoServer.ApplyMiddleware(gwmux),
			}

			return gwServer.ListenAndServe()
		})
	}
	{{- end }}

	return eg.Wait()
}
`
