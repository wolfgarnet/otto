package walk

import (
	"github.com/robertkrimen/otto/ast"
	"fmt"
	"reflect"
)

type Walker2 struct {

}

func (w *Walker2) Walk(node ast.Node) {
	fmt.Printf("WALKING: %v\n", reflect.TypeOf(node))

	switch n := node.(type) {
	case *ast.BinaryExpression:
		w.VisitBinary(n)
	case *ast.ExpressionStatement:
		w.VisitExpression(n)
	case *ast.Identifier:
		w.VisitIdentifier(n)
	case *ast.Program:
		w.VisitProgram(n)
	}
}

func (w *Walker2) VisitProgram(node *ast.Program) {
	println("Visiting program", node)

	for _, e := range node.Body {
		w.Walk(e)
	}

	// Walking function and variable declarations
	for _, value := range node.DeclarationList {
		switch value := value.(type) {
		case *ast.FunctionDeclaration:
			w.Walk(value.Function)
		case *ast.VariableDeclaration:
			for _, value := range value.List {
				w.Walk(value)
			}
		default:
			panic(fmt.Errorf("Here be dragons: cmpl.parseProgram.DeclarationList(%T)", value))
		}
	}
}

func (w *Walker2) VisitBinary(node *ast.BinaryExpression) {
	println("Visiting binary", node)
	w.Walk(node.Left)
	w.Walk(node.Right)
}

func (w *Walker2) VisitExpression(node *ast.ExpressionStatement) {
	println("Visiting expression", node)
	w.Walk(node.Expression)
}

func (w *Walker2) VisitIdentifier(node *ast.Identifier) {
	println("Visiting identifier222222", node)
}
