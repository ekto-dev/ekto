package templates

const RpcClientServiceTpl = `
{{ $svcName := .Name }}
type {{ $svcName }}Querier struct {
	//cc   string
	client {{ $svcName }}Client
}

func New{{ $svcName }}Querier(cc *grpc.ClientConn) *{{ $svcName }}Querier {
	return &{{ $svcName }}Querier{
		client: New{{ $svcName }}Client(cc),
	}
}

//func (q *{{ $svcName }}Querier) newClient(ctx context.Context) ({{ $svcName }}Client, error) {
//	cc, err := grpc.DialContext(ctx, q.addr)
//	if err != nil {
//		return nil, err
//	}
//
//	return New{{ $svcName }}Client(cc), nil
//}

{{- range .Methods }}
{{- if hasQueryMethod . }}
func (q *{{ $svcName }}Querier) {{ queryMethod . }}(ctx context.Context, id string) (*{{ output . }}, error) {
	return q.client.{{ name . }}(ctx, &{{ input . }}{Id: id})
}
{{- end }}
{{- end }}
`
