package templates

const RpcClientTpl = `// Code generated by protoc-gen-ekto. DO NOT EDIT.
// source: {{ .InputPath }}
package {{ package . }}

import (
	"context"
	"net"
	"net/http"

	ektoserver "github.com/ekto-dev/ekto/server"
	ektorpc "github.com/ekto-dev/ekto/server/rpc"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/encoding/protojson"
)
{{ "" }}
func init() {
	{{- range .Services }}
	{{- if hasQueryMethods . }}
	ektorpc.RegisterQuerier(func (cc *grpc.ClientConn) ektorpc.Querier[{{ queryableMessage . }}] {
		return New{{ .Name }}Querier(cc)
	})
	{{- end }}
	{{- end }}
}
{{ "" }}
{{- range .Services }}
{{ template "service" . }}
{{- end }}
`
