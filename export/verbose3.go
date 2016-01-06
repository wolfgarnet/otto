package export

import (
	"bytes"
	"github.com/robertkrimen/otto/ast"
	"github.com/robertkrimen/otto/walk"
	"fmt"
)

type Verbose3 struct {
	walk.Visitor3Impl
	buffer bytes.Buffer
	level int
}

func (v *Verbose3) prependLevel() {
	for i := 0; i < v.level * 3; i++ {
		v.buffer.WriteString(" ")
	}
}

func (v *Verbose3) VisitAssign(w *walk.Walker3, node *ast.AssignExpression, parent ast.Node) {
	w.Walk(node.Left, node)
	v.buffer.WriteString(" ")
	v.buffer.WriteString(node.Operator.String())
	v.buffer.WriteString(" ")
	w.Walk(node.Right, node)
}

func (v *Verbose3) VisitBinary(w *walk.Walker3, node *ast.BinaryExpression, parent ast.Node) {
	v.buffer.WriteString("(")
	w.Walk(node.Left, node)
	v.buffer.WriteString(" ")
	v.buffer.WriteString(node.Operator.String())
	v.buffer.WriteString(" ")
	w.Walk(node.Right, node)
	v.buffer.WriteString(")")
}

/*
func (v *Verbose3) VisitBlock(w *walk.Walker3, node *ast.BlockStatement, parent ast.Node) {
	for _, value := range node.List {
		w.Walk(value, node)
	}
}
*/

func (v *Verbose3) VisitBoolean(w *walk.Walker3, node *ast.BooleanLiteral, parent ast.Node) {
	v.buffer.WriteString(node.Literal)
}

func (v *Verbose3) VisitCall(w *walk.Walker3, node *ast.CallExpression, parent ast.Node) {
	w.Walk(node.Callee, node)
	v.buffer.WriteString("(")
	for _, value := range node.ArgumentList {
		w.Walk(value, node)
	}
	v.buffer.WriteString(")")
}

func (v *Verbose3) VisitExpression(w *walk.Walker3, node *ast.ExpressionStatement, parent ast.Node) {
	v.prependLevel()
	w.Walk(node.Expression, node)
	v.buffer.WriteString(";\n")
}

func (v *Verbose3) VisitFor(w *walk.Walker3, node *ast.ForStatement, parent ast.Node) {
	v.buffer.WriteString("for(")
	w.Walk(node.Initializer, node)
	v.buffer.WriteString(" ; ")
	w.Walk(node.Test, node)
	v.buffer.WriteString(" ; ")
	w.Walk(node.Update, node)
	v.buffer.WriteString(") {\n")
	v.level++
	w.Walk(node.Body, node)
	v.level--
	v.buffer.WriteString("}")
}

func (v *Verbose3) VisitFunction(w *walk.Walker3, node *ast.FunctionLiteral, parent ast.Node) {
	v.buffer.WriteString("function ")
	w.Walk(node.Name, node)
	v.buffer.WriteString("(")
	for _, value := range node.ParameterList.List {
		w.Walk(value, node)
	}
	v.buffer.WriteString(") {\n")
	v.level++
	w.Walk(node.Body, node)
	v.buffer.WriteString("}")

	v.level--

	for _, value := range node.DeclarationList {
		switch value := value.(type) {
		case *ast.FunctionDeclaration:
			w.Walk(value.Function, node)
		case *ast.VariableDeclaration:
			for _, value := range value.List {
				w.Walk(value, node)
			}
		default:
			panic(fmt.Errorf("Here be dragons: parseProgram.declaration(%T)", value))
		}
	}
}

func (v *Verbose3) VisitIdentifier(w *walk.Walker3, node *ast.Identifier, parent ast.Node) {
	v.buffer.WriteString(node.Name)
}

func (v *Verbose3) VisitNumber(w *walk.Walker3, node *ast.NumberLiteral, parent ast.Node) {
	v.buffer.WriteString(node.Literal)
}

/*
func (v *Verbose3) VisitSequence(w *walk.Walker3, node *ast.SequenceExpression, parent ast.Node) {
	println("[DEFAULT] Visiting sequence", node)
	for _, e := range node.Sequence {
		w.Walk(e, node)
		v.buffer.WriteString(";")
	}
}
*/

func (v *Verbose3) VisitString(w *walk.Walker3, node *ast.StringLiteral, parent ast.Node) {
	v.buffer.WriteString(node.Literal)
}

func (v *Verbose3) VisitUnary(w *walk.Walker3, node *ast.UnaryExpression, parent ast.Node) {
	if !node.Postfix {
		v.buffer.WriteString(node.Operator.String())
	}
	w.Walk(node.Operand, node)
	if node.Postfix {
		v.buffer.WriteString(node.Operator.String())
	}
}

func (v *Verbose3) VisitWhile(w *walk.Walker3, node *ast.WhileStatement, parent ast.Node) {
	v.buffer.WriteString("while(")
	w.Walk(node.Test, node)
	v.buffer.WriteString(") {\n")
	v.level++
	w.Walk(node.Body, node)
	v.level--
	v.buffer.WriteString("}")
}


func (v Verbose3) ToString() string {
	return v.buffer.String()
}