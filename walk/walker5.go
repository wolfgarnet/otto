package walk

import "github.com/robertkrimen/otto/ast"

type Visitor5 interface {
	VisitBinary(binary *ast.BinaryExpression)
	VisitExpression(expression *ast.ExpressionStatement)
	VisitIdentifier(identifier *ast.Identifier)
	VisitProgram(program *ast.Program)
}

type Walker5 struct {
	Visitor Visitor5
}

func (w Walker5) Walk(node *ast.Node) {

}
