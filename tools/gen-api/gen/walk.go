package gen

import (
	"reflect"

	"github.com/dop251/goja/ast"
)

// Visitor Enter method is invoked for each node encountered by Walk.
// If the result visitor w is not nil, Walk visits each of the children
// of node with the visitor v, followed by a call of the Exit method.
type Visitor interface {
	Enter(n ast.Node) Visitor
	Exit(n ast.Node)
}

// Walk traverses an AST in depth-first order: It starts by calling
// v.Enter(node); node must not be nil. If the visitor v returned by
// v.Enter(node) is not nil, Walk is invoked recursively with visitor
// v for each of the non-nil children of node, followed by a call
// of v.Exit(node).
func Walk(v Visitor, n ast.Node) {
	if n == nil {
		return
	}

	if v = v.Enter(n); v == nil {
		return
	}

	defer v.Exit(n)

	nType := reflect.TypeOf(n)
	nValue := reflect.ValueOf(n)

	if nType.Kind() == reflect.Ptr {
		nType = nType.Elem()
	}
	if nValue.Kind() == reflect.Ptr {
		nValue = nValue.Elem()
	}

	for i := 0; i < nValue.NumField(); i++ {
		field := nValue.Field(i).Interface()

		if node, ok := field.(ast.Node); ok {
			Walk(v, node)
		}

		if nType.Field(i).Type.Kind() == reflect.Slice {
			for j := 0; j < nValue.Field(i).Len(); j++ {
				fieldValue := nValue.Field(i).Index(j)
				if fieldValue.Kind() == reflect.Struct {
					fieldValue = fieldValue.Addr()
				}
				field := fieldValue.Interface()
				if node, ok := field.(ast.Node); ok {
					Walk(v, node)
				}
			}
		}
	}

}
