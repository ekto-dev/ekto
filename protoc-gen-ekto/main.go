package main

import (
	"github.com/ekto-dev/ekto/protoc-gen-ekto/module"
	pgs "github.com/lyft/protoc-gen-star/v2"
	pgsgo "github.com/lyft/protoc-gen-star/v2/lang/go"
)

func main() {
	pgs.Init(pgs.DebugEnv("DEBUG")).
		RegisterModule(module.Generator()).
		RegisterPostProcessor(pgsgo.GoFmt(), pgsgo.GoImports()).
		Render()
}
