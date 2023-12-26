package module

import (
	"fmt"
	"html/template"

	"github.com/ekto-dev/ekto/protoc-gen-ekto/ekto"
	"github.com/ekto-dev/ekto/protoc-gen-ekto/templates"
	pgs "github.com/lyft/protoc-gen-star/v2"
	pgsgo "github.com/lyft/protoc-gen-star/v2/lang/go"
	"github.com/samber/lo"
	"google.golang.org/genproto/googleapis/api/annotations"
)

type EktoModule struct {
	ctx pgsgo.Context
	*pgs.ModuleBase
	templateFns map[string]any
}

func Generator() *EktoModule {
	return &EktoModule{
		ModuleBase: &pgs.ModuleBase{},
	}
}

func (m *EktoModule) Name() string { return "ekto" }

func (m *EktoModule) InitContext(c pgs.BuildContext) {
	m.ModuleBase.InitContext(c)
	m.ctx = pgsgo.InitContext(c.Parameters())
	m.templateFns = map[string]any{
		"input": func(method pgs.Method) string {
			m.ModuleBase.Debug("input:", method.Input().Name().String())
			if m.ctx.ImportPath(method.Input()) != m.ctx.ImportPath(method) {
				return fmt.Sprintf("%s.%s", m.ctx.PackageName(method.Input()).String(), m.ctx.Name(method.Input()).String())
			}

			return m.ctx.Name(method.Input()).String()
		},
		"messageHandlesEvent": func(msg pgs.Message) bool {
			opts := &ekto.MessageOptions{}
			ok, err := msg.Extension(ekto.E_Msg, opts)
			if !ok || err != nil {
				return false
			}

			return opts.Mq != nil && opts.Mq.EventName != ""
		},
		"messageEventName": func(msg pgs.Message) string {
			opts := &ekto.MessageOptions{}
			ok, err := msg.Extension(ekto.E_Msg, opts)
			if !ok || err != nil {
				return ""
			}

			if opts.Mq == nil {
				return ""
			}

			return opts.Mq.EventName
		},
	}
}

func (m *EktoModule) Execute(targets map[string]pgs.File, pkgs map[string]pgs.Package) []pgs.Artifact {
	for _, f := range targets {
		m.generateMQFile(f)
		m.generateDatabaseFile(f)
		m.generateServerFile(f)
		m.generateRpcClientFile(f)
	}

	return m.Artifacts()
}

func (m *EktoModule) generateMQFile(f pgs.File) {
	m.Push(f.Name().String())
	defer m.Pop()
	out := m.ctx.OutputPath(f).SetExt(".ekto.mq.go")
	m.ModuleBase.Log("generating mq file", f.Name().String())

	handlesEvent := func(method pgs.Method) bool {
		m.ModuleBase.Debug("handlesEvent:", method.Name().String())
		var ektoMsgOptions = &ekto.MessageOptions{}
		if inputHasEktoConfig, err := method.Input().Extension(ekto.E_Msg, ektoMsgOptions); !inputHasEktoConfig || err != nil {
			return false
		}

		if ektoMsgOptions.Mq != nil && ektoMsgOptions.Mq.EventName != "" {
			return true
		}

		return false
		//var ektoOptions = &ekto.Options{}
		//if defined, _ := method.Extension(ekto.E_Dev, ektoOptions); !defined {
		//	return false // no ekto options defined
		//}
		//
		//return ektoOptions.Mq != nil && ektoOptions.Mq.Handles != ""
	}

	svcHandlesEvent := func(svc pgs.Service) bool {
		for _, method := range svc.Methods() {
			if handlesEvent(method) {
				return true
			}
		}

		return false
	}
	tpl := template.New("ekto-mq").Funcs(map[string]any{
		"package": m.ctx.PackageName,
		"name":    m.ctx.Name,
		"input":   m.templateFns["input"],
		"output": func(m pgs.Method) string {
			return m.Output().Name().String()
		},
		"hasMessageHandler": svcHandlesEvent,
		"handlesEvent":      handlesEvent,
		"eventName": func(method pgs.Method) string {
			var ektoMsgOptions = &ekto.MessageOptions{}
			if inputHasEktoConfig, err := method.Input().Extension(ekto.E_Msg, ektoMsgOptions); !inputHasEktoConfig || err != nil {
				return ""
			}

			if ektoMsgOptions.Mq != nil && ektoMsgOptions.Mq.EventName != "" {
				return ektoMsgOptions.Mq.EventName
			}

			return ""
		},
		"messageHandlesEvent": m.templateFns["messageHandlesEvent"],
		"messageEventName":    m.templateFns["messageEventName"],
	})

	hasHandlers := false
	for _, msg := range f.Messages() {
		if m.templateFns["messageHandlesEvent"].(func(m pgs.Message) bool)(msg) {
			hasHandlers = true
			break
		}
	}
	if !hasHandlers {
		for _, svc := range f.Services() {
			if svcHandlesEvent(svc) {
				hasHandlers = true
				break
			}
		}
	}

	if !hasHandlers {
		return
	}

	template.Must(tpl.Parse(templates.MqTpl))
	template.Must(tpl.New("service").Parse(templates.MQServiceTpl))
	m.AddGeneratorTemplateFile(out.String(), tpl, f)
}

func (m *EktoModule) generateServerFile(f pgs.File) {
	m.Push(f.Name().String())
	defer m.Pop()
	out := m.ctx.OutputPath(f).SetExt(".ekto.server.go")

	tpl := template.New("ekto-server").Funcs(map[string]any{
		"package": m.ctx.PackageName,
		"name":    m.ctx.Name,
		"hasGateway": func(svc pgs.Service) bool {
			for _, method := range svc.Methods() {
				e := &annotations.Http{}
				if defined, _ := method.Extension(annotations.E_Http, e); defined {
					return true
				}
			}

			return false
		},
	})
	template.Must(tpl.Parse(templates.ServerTpl))
	template.Must(tpl.New("service").Parse(templates.ServerServiceTpl))
	m.AddGeneratorTemplateFile(out.String(), tpl, f)
}

func (m *EktoModule) generateRpcClientFile(f pgs.File) {
	m.Push(f.Name().String())
	defer m.Pop()
	out := m.ctx.OutputPath(f).SetExt(".ekto.client.go")
	m.Debug("generateRpcClientFile:", out.String())

	queryableMessage, found := lo.Find(f.AllMessages(), func(m pgs.Message) bool {
		e := &ekto.MessageOptions{}
		if defined, _ := m.Extension(ekto.E_Msg, e); defined && e.Queryable != false {
			return true
		}
		return false
	})
	queryableMessageName := ""
	if found {
		queryableMessageName = queryableMessage.Name().String()
	}

	tpl := template.New("ekto-client").Funcs(map[string]any{
		"package":          m.ctx.PackageName,
		"name":             m.ctx.Name,
		"queryableMessage": func(_ pgs.Node) string { return queryableMessageName },
		"queryableMessageFQN": func(_ pgs.Node) string {
			if found {
				return queryableMessage.FullyQualifiedName()
			}
			return ""
		},
		"hasQueryMethods": func(svc pgs.Service) bool {
			for _, method := range svc.Methods() {
				e := &ekto.Options{}
				if defined, _ := method.Extension(ekto.E_Dev, e); defined && e.Querier != nil {
					return true
				}
			}

			return false
		},
		"queryMethod": func(method pgs.Method) string {
			e := &ekto.Options{}
			if defined, _ := method.Extension(ekto.E_Dev, e); defined && e.Querier != nil {
				switch e.Querier.Method {
				case ekto.QuerierMethod_FIND:
					return "Find"
				case ekto.QuerierMethod_LIST:
					return "List"
				case ekto.QuerierMethod_CREATE:
					return "Create"
				case ekto.QuerierMethod_UPDATE:
					return "Update"
				case ekto.QuerierMethod_DELETE:
					return "Delete"
				default:
					return ""
				}
			}

			return ""
		},
		"hasQueryMethod": func(method pgs.Method) bool {
			e := &ekto.Options{}
			if defined, _ := method.Extension(ekto.E_Dev, e); defined && e.Querier != nil {
				return true
			}

			return false
		},
		"input": m.templateFns["input"],
		"output": func(method pgs.Method) string {
			return method.Output().Name().String()
		},
	})
	template.Must(tpl.Parse(templates.RpcClientTpl))
	template.Must(tpl.New("service").Parse(templates.RpcClientServiceTpl))
	m.AddGeneratorTemplateFile(out.String(), tpl, f)
}

func (m *EktoModule) generateDatabaseFile(f pgs.File) {
	m.Push(f.Name().String())
	defer m.Pop()
	out := m.ctx.OutputPath(f).SetExt(".ekto.db.go")

	tpl := template.New("ekto-db").Funcs(map[string]any{
		"package": m.ctx.PackageName,
		"name":    m.ctx.Name,
		"hasDatabase": func(svc pgs.Service) bool {
			var ektoOptions = &ekto.SvcOptions{}
			if defined, _ := svc.Extension(ekto.E_Svc, ektoOptions); defined {
				if ektoOptions.Db != nil && ektoOptions.Db.Name != "" {
					return true
				}
			}

			return false
		},
		"databaseName": func(svc pgs.Service) string {
			var ektoOptions = &ekto.SvcOptions{}
			if defined, _ := svc.Extension(ekto.E_Svc, ektoOptions); defined {
				if ektoOptions.Db != nil && ektoOptions.Db.Name != "" {
					return ektoOptions.Db.Name
				}
			}

			return ""
		},
	})
	template.Must(tpl.Parse(templates.DbTpl))
	template.Must(tpl.New("connect").Parse(templates.DbConnectTpl))
	m.AddGeneratorTemplateFile(out.String(), tpl, f)
}

var _ pgs.Module = (*EktoModule)(nil)
