package gen

import (
	"encoding/json"
	"os"

	"github.com/dop251/goja/ast"
	"github.com/dop251/goja/parser"
	"github.com/tdewolff/parse/v2"
	"github.com/tdewolff/parse/v2/js"
)

func ParseSchemaFilex(f string) ([]ApiSchema, error) {

	apiSchemaContent, err := os.ReadFile(f)
	if err != nil {
		return nil, err
	}

	schemaJsAST, err := js.Parse(parse.NewInputBytes(apiSchemaContent), js.Options{})
	if err != nil {
		return nil, err
	}

	var schemaJson string
	js.Walk(apiSchemaWalkx(&schemaJson), schemaJsAST)

	os.WriteFile("x.json", []byte(schemaJson), 0644)

	result := []ApiSchema{}
	err = json.Unmarshal([]byte(schemaJson), &result)
	if err != nil {
		return nil, err
	}

	return result, err
}

func apiSchemaWalkx(apiSchema *string) walkerFuncx {
	return func(s string) {
		*apiSchema = s
	}
}

type walkerFuncx func(string)

func (w walkerFuncx) Enter(n js.INode) js.IVisitor {
	switch n := n.(type) {
	case *js.BindingElement:
		if n.Binding.String() == "apiSchema" {
			w(n.Default.JS())
		}
	}

	return w
}
func (w walkerFuncx) Exit(n js.INode) {
}

func ParseSchemaFile(f string) ([]ApiSchema, error) {

	schemaJsAST, err := parser.ParseFile(nil, f, nil, 0)
	if err != nil {
		return nil, err
	}

	var startPos, endPos int
	Walk(apiSchemaIdxWalk(&startPos, &endPos), schemaJsAST)
	schemaJson := schemaJsAST.File.Source()[startPos-1 : endPos]

	result := []ApiSchema{}
	err = json.Unmarshal([]byte(schemaJson), &result)
	if err != nil {
		return nil, err
	}

	return result, err
}

func apiSchemaIdxWalk(start *int, end *int) walkerFunc {
	return func(startPos, endPos int) {
		*start = startPos
		*end = endPos
	}
}

type walkerFunc func(int, int)

func (w walkerFunc) Enter(n ast.Node) Visitor {
	switch n := n.(type) {
	case *ast.Binding:
		if v, ok := n.Target.(*ast.Identifier); ok {
			if v.Name == "apiSchema" {
				w(int(n.Initializer.Idx0()), int(n.Initializer.Idx1()))
			}
		}
	}

	return w
}
func (w walkerFunc) Exit(n ast.Node) {
}
