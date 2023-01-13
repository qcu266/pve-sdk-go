package gen

import (
	"fmt"

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

	switch n := n.(type) {
	case *ast.ArrayLiteral:
		if n != nil {
			for _, ex := range n.Value {
				Walk(v, ex)
			}
		}
	case *ast.ArrayPattern:
	case *ast.ArrowFunctionLiteral:
	case *ast.AssignExpression:
		if n != nil {
			Walk(v, n.Left)
			Walk(v, n.Right)
		}
	case *ast.AwaitExpression:
	case *ast.BadExpression:
	case *ast.BadStatement:
	case *ast.BinaryExpression:
		if n != nil {
			Walk(v, n.Left)
			Walk(v, n.Right)
		}
	case *ast.Binding:
		if n != nil {
			Walk(v, n.Target)
			Walk(v, n.Initializer)
		}
	case *ast.BlockStatement:
		if n != nil {
			for _, s := range n.List {
				Walk(v, s)
			}
		}
	case *ast.BooleanLiteral:
	case *ast.BracketExpression:
		if n != nil {
			Walk(v, n.Left)
			Walk(v, n.Member)
		}
	case *ast.BranchStatement:
		if n != nil {
			Walk(v, n.Label)
		}
	case *ast.CallExpression:
		if n != nil {
			Walk(v, n.Callee)
			for _, a := range n.ArgumentList {
				Walk(v, a)
			}
		}
	case *ast.CaseStatement:
		if n != nil {
			Walk(v, n.Test)
			for _, c := range n.Consequent {
				Walk(v, c)
			}
		}
	case *ast.CatchStatement:
		if n != nil {
			Walk(v, n.Parameter)
			Walk(v, n.Body)
		}
	case *ast.ClassDeclaration:
	case *ast.ClassLiteral:
	case *ast.ConditionalExpression:
		if n != nil {
			Walk(v, n.Test)
			Walk(v, n.Consequent)
			Walk(v, n.Alternate)
		}
	case *ast.DebuggerStatement:
	case *ast.DotExpression:
		if n != nil {
			Walk(v, n.Left)
			Walk(v, &n.Identifier)
		}
	case *ast.DoWhileStatement:
		if n != nil {
			Walk(v, n.Test)
			Walk(v, n.Body)
		}
	case *ast.EmptyStatement:
	case *ast.ExpressionBody:
	case *ast.ExpressionStatement:
		if n != nil {
			Walk(v, n.Expression)
		}
	case *ast.ForInStatement:
		if n != nil {
			Walk(v, n.Into)
			Walk(v, n.Source)
			Walk(v, n.Body)
		}
	case *ast.ForLoopInitializerExpression:
	case *ast.ForLoopInitializerLexicalDecl:
	case *ast.ForLoopInitializerVarDeclList:
	case *ast.ForOfStatement:
	case *ast.ForStatement:
		if n != nil {
			Walk(v, n.Initializer)
			Walk(v, n.Update)
			Walk(v, n.Test)
			Walk(v, n.Body)
		}
	case *ast.FunctionDeclaration:
		if n != nil {
			Walk(v, n.Function)
		}
	case *ast.FunctionLiteral:
		if n != nil {
			Walk(v, n.Name)
			for _, p := range n.ParameterList.List {
				Walk(v, p)
			}
			Walk(v, n.Body)
		}
	case *ast.Identifier:
	case *ast.IfStatement:
		if n != nil {
			Walk(v, n.Test)
			Walk(v, n.Consequent)
			Walk(v, n.Alternate)
		}
	case *ast.LabelledStatement:
		if n != nil {
			Walk(v, n.Label)
			Walk(v, n.Statement)
		}
	case *ast.LexicalDeclaration:
		for _, p := range n.List {
			Walk(v, p)
		}
	case *ast.MetaProperty:
	case *ast.NewExpression:
		if n != nil {
			Walk(v, n.Callee)
			for _, a := range n.ArgumentList {
				Walk(v, a)
			}
		}
	case *ast.NullLiteral:
	case *ast.NumberLiteral:
	case *ast.ObjectLiteral:
		if n != nil {
			for _, p := range n.Value {
				Walk(v, p)
			}
		}
	case *ast.ObjectPattern:
	case *ast.ParameterList:
	case *ast.PrivateDotExpression:
	case *ast.Program:
		if n != nil {
			for _, b := range n.Body {
				Walk(v, b)
			}
		}
	case *ast.PropertyKeyed:
	case *ast.PropertyShort:
	case *ast.RegExpLiteral:
	case *ast.ReturnStatement:
		if n != nil {
			Walk(v, n.Argument)
		}
	case *ast.SequenceExpression:
		if n != nil {
			for _, e := range n.Sequence {
				Walk(v, e)
			}
		}
	case *ast.StringLiteral:
	case *ast.SuperExpression:
	case *ast.SwitchStatement:
		if n != nil {
			Walk(v, n.Discriminant)
			for _, c := range n.Body {
				Walk(v, c)
			}
		}
	case *ast.TemplateElement:
	case *ast.TemplateLiteral:
	case *ast.ThisExpression:
	case *ast.ThrowStatement:
		if n != nil {
			Walk(v, n.Argument)
		}
	case *ast.TryStatement:
		if n != nil {
			Walk(v, n.Body)
			Walk(v, n.Catch)
			Walk(v, n.Finally)
		}
	case *ast.UnaryExpression:
		if n != nil {
			Walk(v, n.Operand)
		}
	case *ast.VariableStatement:
		if n != nil {
			for _, e := range n.List {
				Walk(v, e)
			}
		}
	case *ast.WhileStatement:
		if n != nil {
			Walk(v, n.Test)
			Walk(v, n.Body)
		}
	case *ast.WithStatement:
		if n != nil {
			Walk(v, n.Object)
			Walk(v, n.Body)
		}
	default:
		panic(fmt.Sprintf("Walk: unexpected node type %T", n))
	}
}
