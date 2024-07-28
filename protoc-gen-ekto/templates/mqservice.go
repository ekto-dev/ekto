package templates

const MQServiceTpl = `type {{ .Name }}MQProxy struct {
	server *grpc.Server
}

func New{{ .Name }}MQProxy() *{{ .Name }}MQProxy {
	return &{{ .Name }}MQProxy{
		server: grpc.NewServer(),
	}
}

func (p *{{ .Name }}MQProxy) Register(svc {{ name . }}) {
	Register{{ name . }}(p.server, svc)
}

func (p *{{ .Name }}MQProxy) Run(ctx context.Context, clientBuilder func (ctx context.Context, topic string, handlerName string) (cloudeventsv2.Client, error)) error {
	// Start the gRPC server in a goroutine
	var port int
	var wgForPort sync.WaitGroup
	wgForPort.Add(1)
	go func() {
		lis, err := net.Listen("tcp", ":0")

		if err != nil {
			log.Fatalf("failed to listen: %s", err)
		}
		port = lis.Addr().(*net.TCPAddr).Port
		wgForPort.Done()

		if err := p.server.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %s", err)
		}
	}()
	wgForPort.Wait()

	// connect to the gRPC server
	conn, err := grpc.DialContext(
		ctx,
		fmt.Sprintf(":%d", port),
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		return err
	}

	svcClient := New{{ .Name }}Client(conn)
	eg, ctx := errgroup.WithContext(ctx)

	// For each event, start a new client, and start a receiver
	{{- range .Methods }}
	{{- if handlesEvent . }}
	{
		client, err := clientBuilder(ctx, "{{ eventName . }}", "{{ .Name }}")
		if err != nil {
			return err
		}

		eg.Go(func() error {
			return client.StartReceiver(ctx, func(ctx context.Context, event cloudeventsv2.Event) protocol.Result {
				// decode the cloudevent into a protobuf message
				protoEvent, err := format.ToProto(&event)
				if err != nil {
					return protocol.NewReceipt(false, "failed to decode event: %s", err.Error())
				}
	
				msg := &{{ input . }}{}
				err = protoEvent.GetProtoData().UnmarshalTo(msg)
				if err != nil {
					return protocol.NewReceipt(false, "failed to unmarshal event data: %s", err.Error())
				}
	
				_, err = svcClient.{{ .Name }}(ctx, msg)
				if err != nil {
					return protocol.NewReceipt(false, "failed to call service method: %s", err.Error())
				}

				return protocol.NewReceipt(true, "")
			})
		})
	}

	{{- end }}
	{{- end }}
	return eg.Wait()
}
`
