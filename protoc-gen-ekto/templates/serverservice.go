package templates

const ServerServiceTpl = `
func Start{{ .Name }}Server(ctx context.Context, rpcListenAddr string, srv {{ name . }}, opts ...ektoserver.Option) error {
	ektoServer := ektoserver.NewEktoServer(opts...)
	server := grpc.NewServer()
	Register{{ name . }}(server, srv)

	eg, _ := errgroup.WithContext(ctx)
	eg.Go(func() error {
		lis, err := net.Listen("tcp", rpcListenAddr)

		if err != nil {
			return err
		}

		return server.Serve(lis)
	})

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
				runtime.WithMarshalerOption(
					runtime.MIMEWildcard,
					&runtime.JSONPb{
						MarshalOptions: protojson.MarshalOptions{
							UseProtoNames: true,
						},
					},
				),
			)

			err = Register{{ .Name }}Handler(ctx, gwmux, conn)
			if err != nil {
				return err
			}

			gwServer := &http.Server{
				Addr:    ektoServer.GatewayAddr(),
				Handler: gwmux,
			}

			return gwServer.ListenAndServe()
		})
	}

	return eg.Wait()
}
`