package walk

import (
	"github.com/robertkrimen/otto/ast"
	"fmt"
	"reflect"
)

type Walker struct {
}

type Visitor interface {
	Visit(node ast.Node) (w Visitor)
}

func (w Walker) Walk(v Visitor, node ast.Node) {
	if v = v.Visit(node); v == nil {
		return
	}

	switch t := node.(type) {
	case *ast.ArrayLiteral:
		fmt.Printf("Walking %v\n", reflect.TypeOf(t))
		for _, e := range t.Value {
			w.Walk(v, e)
		}

	case *ast.AssignExpression:
		fmt.Printf("Walking %v\n", reflect.TypeOf(t))
		w.Walk(v, t.Left)
		w.Walk(v, t.Right)

	case *ast.BadExpression:
		fmt.Printf("Walking %v\n", reflect.TypeOf(t))

	case *ast.BadStatement:
		fmt.Printf("Walking %v\n", reflect.TypeOf(t))

	case *ast.BinaryExpression:
		fmt.Printf("Walking %v\n", reflect.TypeOf(t))
		//w.Walk(v, t)
		w.Walk(v, t.Left)
		w.Walk(v, t.Right)

	case *ast.BlockStatement:
		fmt.Printf("Walking %v\n", reflect.TypeOf(t))
		for _, e := range t.List {
			w.Walk(v, e)
		}

	case *ast.BooleanLiteral:
		fmt.Printf("Walking %v\n", reflect.TypeOf(t))

	case *ast.BracketExpression:
		fmt.Printf("Walking %v\n", reflect.TypeOf(t))
		w.Walk(v, t.Left)
		w.Walk(v, t.Member)

	case *ast.BranchStatement:
		fmt.Printf("Walking %v\n", reflect.TypeOf(t))
		w.Walk(v, t.Label)

	case *ast.CallExpression:
		fmt.Printf("Walking %v\n", reflect.TypeOf(t))
		w.Walk(v, t.Callee)
		for _, e := range t.ArgumentList {
			w.Walk(v, e)
		}

	case *ast.CaseStatement:
		fmt.Printf("Walking %v\n", reflect.TypeOf(t))
		w.Walk(v, t.Test)
		for _, e := range t.Consequent {
			w.Walk(v, e)
		}

	case *ast.CatchStatement:
		fmt.Printf("Walking %v\n", reflect.TypeOf(t))
		w.Walk(v, t.Parameter)
		w.Walk(v, t.Body)

	case *ast.ConditionalExpression:
		fmt.Printf("Walking %v\n", reflect.TypeOf(t))
		w.Walk(v, t.Test)
		w.Walk(v, t.Consequent)
		w.Walk(v, t.Alternate)

	case *ast.DebuggerStatement:
		fmt.Printf("Walking %v\n", reflect.TypeOf(t))

	case *ast.DotExpression:
		fmt.Printf("Walking %v\n", reflect.TypeOf(t))
		w.Walk(v, t.Left)
		w.Walk(v, &t.Identifier)

	case *ast.DoWhileStatement:
		fmt.Printf("Walking %v\n", reflect.TypeOf(t))
		w.Walk(v, t.Test)
		w.Walk(v, t.Body)

	case *ast.EmptyStatement:
		fmt.Printf("Walking %v\n", reflect.TypeOf(t))

	case *ast.ExpressionStatement:
		fmt.Printf("Walking %v\n", reflect.TypeOf(t))
		w.Walk(v, t.Expression)

	case *ast.ForInStatement:
		fmt.Printf("Walking %v\n", reflect.TypeOf(t))
		w.Walk(v, t.Into)
		w.Walk(v, t.Body)

	case *ast.ForStatement:
		fmt.Printf("Walking %v\n", reflect.TypeOf(t))
		w.Walk(v, t.Initializer)
		w.Walk(v, t.Test)
		w.Walk(v, t.Update)
		w.Walk(v, t.Body)

	case *ast.FunctionLiteral:
		fmt.Printf("Walking %v\n", reflect.TypeOf(t))
		w.Walk(v, t.Name)
		// w.Walk(t.ParameterList) ?
		w.Walk(v, t.Body)
		for _, value := range t.DeclarationList {
			switch value := value.(type) {
			case *ast.FunctionDeclaration:
				w.Walk(v, value.Function)
			case *ast.VariableDeclaration:
				for _, value := range value.List {
					w.Walk(v, value)
				}
			default:
				panic(fmt.Errorf("Here be dragons: parseProgram.declaration(%T)", value))
			}
		}

	case *ast.FunctionStatement:
		fmt.Printf("Walking %v\n", reflect.TypeOf(t))
		w.Walk(v, t.Function)

	case *ast.Identifier:
		fmt.Printf("Walking %v\n", reflect.TypeOf(t))

	case *ast.IfStatement:
		fmt.Printf("Walking %v\n", reflect.TypeOf(t))
		w.Walk(v, t.Test)
		w.Walk(v, t.Consequent)
		w.Walk(v, t.Alternate)

	case *ast.LabelledStatement:
		fmt.Printf("Walking %v\n", reflect.TypeOf(t))
		w.Walk(v, t.Label)
		w.Walk(v, t.Statement)

	case *ast.NewExpression:
		fmt.Printf("Walking %v\n", reflect.TypeOf(t))
		w.Walk(v, t.Callee)
		for _, e := range t.ArgumentList {
			w.Walk(v, e)
		}

	case *ast.NullLiteral:
		fmt.Printf("Walking %v\n", reflect.TypeOf(t))

	case *ast.NumberLiteral:
		fmt.Printf("Walking %v\n", reflect.TypeOf(t))

	case *ast.ObjectLiteral:
		fmt.Printf("Walking %v\n", reflect.TypeOf(t))

/*	case *ast.ParameterList:
		fmt.Printf("Walking %v\n", reflect.TypeOf(t))*/
/*	case *ast.Property:
		fmt.Printf("Walking %v\n", reflect.TypeOf(t))*/
	case *ast.ReturnStatement:
		fmt.Printf("Walking %v\n", reflect.TypeOf(t))
		w.Walk(v, t.Argument)

	case *ast.Program:
		fmt.Printf("Walking %v\n", reflect.TypeOf(t))
		for _, e := range t.Body {
			w.Walk(v, e)
		}

		// Walking function and variable declarations
		for _, value := range t.DeclarationList {
			switch value := value.(type) {
			case *ast.FunctionDeclaration:
				w.Walk(v, value.Function)
			case *ast.VariableDeclaration:
				for _, value := range value.List {
					w.Walk(v, value)
				}
			default:
				panic(fmt.Errorf("Here be dragons: cmpl.parseProgram.DeclarationList(%T)", value))
			}
		}

	case *ast.RegExpLiteral:
		fmt.Printf("Walking %v\n", reflect.TypeOf(t))

	case *ast.SequenceExpression:
		fmt.Printf("Walking %v\n", reflect.TypeOf(t))
		for _, e := range t.Sequence {
			w.Walk(v, e)
		}

	case *ast.StringLiteral:
		fmt.Printf("Walking %v\n", reflect.TypeOf(t))

	case *ast.SwitchStatement:
		fmt.Printf("Walking %v\n", reflect.TypeOf(t))
		w.Walk(v, t.Discriminant)
		for _, e := range t.Body {
			w.Walk(v, e)
		}

	case *ast.ThisExpression:
		fmt.Printf("Walking %v\n", reflect.TypeOf(t))

	case *ast.ThrowStatement:
		fmt.Printf("Walking %v\n", reflect.TypeOf(t))
		w.Walk(v, t.Argument)

	case *ast.TryStatement:
		fmt.Printf("Walking %v\n", reflect.TypeOf(t))
		w.Walk(v, t.Body)
		w.Walk(v, t.Catch)
		w.Walk(v, t.Finally)

	case *ast.UnaryExpression:
		fmt.Printf("Walking %v\n", reflect.TypeOf(t))
		w.Walk(v, t.Operand)

	case *ast.VariableExpression:
		fmt.Printf("Walking %v\n", reflect.TypeOf(t))
		w.Walk(v, t.Initializer)

	case *ast.VariableStatement:
		fmt.Printf("Walking %v\n", reflect.TypeOf(t))
		for _, e := range t.List {
			w.Walk(v, e)
		}

	case *ast.WhileStatement:
		fmt.Printf("Walking %v\n", reflect.TypeOf(t))
		w.Walk(v, t.Test)
		w.Walk(v, t.Body)

	case *ast.WithStatement:
		fmt.Printf("Walking %v\n", reflect.TypeOf(t))
		w.Walk(v, t.Object)
		w.Walk(v, t.Body)

	default:
		fmt.Printf("Waaat!? %v\n", reflect.TypeOf(t))
	}
}
