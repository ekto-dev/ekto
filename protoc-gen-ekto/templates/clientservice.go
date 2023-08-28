package templates

const ClientServiceTpl = `
// New{{ .Name }}ClientConn creates a *grpc.ClientConn to {{ serviceName . }}:50051.
// Unlike New{{ .Name }}ConnectedClient, it does not default to a blocking connection,
// meaning you cannot use the cancel context by default.
func New{{ .Name }}ClientConn(ctx context.Context, opts ...grpc.DialOption) (*grpc.ClientConn, error) {
	return grpc.DialContext(ctx, "{{ serviceName . }}:50051", opts...)
}

// New{{ .Name }}ConnectedClient creates a gRPC client with a
// client connection on the specified name, to port 50051.
// This will always set the blocking option, which will
// allow you to close the connection by cancelling the
// context. For more control, you may call
// New{{ .Name }}ClientConn and pass the resulting
// *grpc.ClientConn into New{{ .Name }}Client.
func New{{ .Name }}ConnectedClient(ctx context.Context, opts ...grpc.DialOption) ({{ .Name }}Client, error) {
	opts = append(opts, grpc.WithBlock())
	cc, err := New{{ .Name }}ClientConn(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return New{{ .Name }}Client(cc), nil
}
`
