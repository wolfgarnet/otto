package walk

import "github.com/robertkrimen/otto/ast"

type Visitor4 interface {
	VisitBinary(binary *ast.BinaryExpression)
	VisitExpression(expression *ast.ExpressionStatement)
	VisitIdentifier(identifier *ast.Identifier)
	VisitProgram(program *ast.Program)
}

func Walk4(v Visitor4, node ast.Node) {
	if v == nil {
		return
	}

	switch n := node.(type) {
	case *ast.BinaryExpression:
		v.VisitBinary(n)
	case *ast.ExpressionStatement:
		v.VisitExpression(n)
	case *ast.Identifier:
		v.VisitIdentifier(n)
	case *ast.Program:
		v.VisitProgram(n)
	}
}
//
//
//type Visitor4Impl struct {
//
//}
//
//func (v *Visitor4Impl) VisitProgram(node *ast.Program) {
//	println("Visiting program", node)
//
//	for _, e := range node.Body {
//		Walk4(e)
//	}
//
//	// Walking function and variable declarations
//	for _, value := range node.DeclarationList {
//		switch value := value.(type) {
//		case *ast.FunctionDeclaration:
//			w.Walk(v, value.Function)
//		case *ast.VariableDeclaration:
//			for _, value := range value.List {
//				w.Walk(v, value)
//			}
//		default:
//			panic(fmt.Errorf("Here be dragons: cmpl.parseProgram.DeclarationList(%T)", value))
//		}
//	}
//}
//
//func (v *Visitor4Impl) VisitBinary(node *ast.BinaryExpression) {
//	println("Visiting binary", node)
//	w.Walk(v, node.Left)
//	w.Walk(v, node.Right)
//}
//
//func (v *Visitor4Impl) VisitExpression(node *ast.ExpressionStatement) {
//	println("Visiting expression", node)
//	w.Walk(v, node.Expression)
//}
//
//func (v *Visitor4Impl) VisitIdentifier(node *ast.Identifier) {
//	println("Visiting identifier", node)
//}
//
