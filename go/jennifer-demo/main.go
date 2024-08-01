package main

import (
	"fmt"

	"github.com/dave/jennifer/jen"
)

func main() {
	f := jen.NewFile("ci")

	f.ImportName("app/usecase/ui", "ui")

	f.Type().Id("testGenerateImp").Struct()

	f.Func().Id("NewTestGenerate").Params().Qual(
		"app/controllers/ci", "TestGenerate",
	).Block(
		jen.Return(jen.Op("&").Id("testGenerateImp").Values()),
	)

	fmt.Printf("%#v", f)
}
