package templates

const ServiceTpl = `type {{ .Name }}MQProxy struct {
	server *grpc.Server
}

func New{{ .Name }}MQProxy() *{{ .Name }}MQProxy {
	return &{{ name . }}MQProxy{
		server: grpc.NewServer(),
	}
}

func (p *{{ .Name }}MQProxy) Register(svc {{ name . }}) {
	Register{{ name . }}(p.server, svc)
}

func (p *{{ .Name }}MQProxy) Run(ctx context.Context, client *cloudeventsv2.Client) error {
	// Start the gRPC server in a goroutine
	go func() {
		lis, err := net.Listen("tcp", ektoPort)

		if err != nil {
			log.Fatalf("failed to listen: %s", err)
		}

		if err := p.server.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %s", err)
		}
	}()

	// connect to the gRPC server
	conn, err := grpc.Dial(
		ctx,
		ektoPort,
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	svcClient := New{{ .Name }}Client(conn)

	return client.StartReceiver(ctx, func(ctx context.Context, event cloudeventsv2.Event) protocol.Result {
		switch event.Type() {
		{{- range .Methods }}
		{{- if handlesEvent . }}
		case "{{ eventName . }}":
			// decode the cloudevent into a protobuf message
			protoEvent, err := format.ToProto(&event)
			if err != nil {
				return protocol.NewReceipt(false, "failed to decode event: %s", err.Error())
			}

			var msg *{{ input . }}
			err = protoEvent.GetProtoData().UnmarshalTo(msg)
			if err != nil {
				return protocol.NewReceipt(false, "failed to unmarshal event data: %s", err.Error())
			}

			_, err = svcClient.{{ .Name }}(ctx, msg)
			if err != nil {
				return protocol.NewReceipt(false, "failed to call service method: %s", err.Error())
			}
			{{- end }}
			{{- end }}
		}

		return protocol.NewReceipt(true, "")
	})
}
`
