package module

import (
	"github.com/ekto-dev/ekto/protoc-gen-ekto/ekto"
	"github.com/ekto-dev/ekto/protoc-gen-ekto/templates"
	pgs "github.com/lyft/protoc-gen-star/v2"
	pgsgo "github.com/lyft/protoc-gen-star/v2/lang/go"
	"html/template"
)

type EktoModule struct {
	ctx pgsgo.Context
	*pgs.ModuleBase
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
}

func (m *EktoModule) Execute(targets map[string]pgs.File, pkgs map[string]pgs.Package) []pgs.Artifact {
	for _, f := range targets {
		m.generateMQFile(f)
		m.generateServerFile(f)
	}

	return m.Artifacts()
}

func (m *EktoModule) generateMQFile(f pgs.File) {
	m.Push(f.Name().String())
	defer m.Pop()
	out := m.ctx.OutputPath(f).SetExt(".ekto.mq.go")

	tpl := template.New("ekto-mq").Funcs(map[string]any{
		"package": m.ctx.PackageName,
		"name":    m.ctx.Name,
		"input": func(method pgs.Method) string {
			m.ModuleBase.Debug("input:", method.Input().Name().String())
			return method.Input().Name().String()
		},
		"output": func(m pgs.Method) string {
			return m.Output().Name().String()
		},
		"handlesEvent": func(method pgs.Method) bool {
			m.ModuleBase.Debug("handlesEvent:", method.Name().String())
			var ektoOptions = &ekto.Options{}
			if defined, _ := method.Extension(ekto.E_Dev, ektoOptions); !defined {
				return false // no ekto options defined
			}

			return ektoOptions.Mq != nil && ektoOptions.Mq.Handles != ""
		},
		"eventName": func(method pgs.Method) string {
			var ektoOptions = &ekto.Options{}
			if defined, _ := method.Extension(ekto.E_Dev, ektoOptions); !defined {
				return "" // no ekto options defined
			}

			return ektoOptions.Mq.Handles
		},
	})
	template.Must(tpl.Parse(templates.MqTpl))
	template.Must(tpl.New("service").Parse(templates.ServiceTpl))
	m.AddGeneratorTemplateFile(out.String(), tpl, f)
}

func (m *EktoModule) generateServerFile(f pgs.File) {
	m.Push(f.Name().String())
	defer m.Pop()
	out := m.ctx.OutputPath(f).SetExt(".ekto.server.go")

	tpl := template.New("ekto-mq").Funcs(map[string]any{
		"package": m.ctx.PackageName,
		"name":    m.ctx.Name,
	})
	template.Must(tpl.Parse(templates.MqTpl))
	template.Must(tpl.New("service").Parse(templates.ServiceTpl))
	m.AddGeneratorTemplateFile(out.String(), tpl, f)
}

var _ pgs.Module = (*EktoModule)(nil)
