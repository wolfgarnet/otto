package export

import (
	"bytes"
	"github.com/robertkrimen/otto/ast"
	"github.com/robertkrimen/otto/walk"
	"fmt"
)

type Verbose struct {
	walk.VisitorImpl
	buffer bytes.Buffer
	level int
}

func (v *Verbose) prependLevel() {
	for i := 0; i < v.level * 3; i++ {
		v.buffer.WriteString(" ")
	}
}

func (v *Verbose) VisitAssign(w *walk.Walker, node *ast.AssignExpression, metadata []walk.Metadata) {
	w.Walk(node.Left, metadata)
	v.buffer.WriteString(" ")
	v.buffer.WriteString(node.Operator.String())
	v.buffer.WriteString(" ")
	w.Walk(node.Right, metadata)
}

func (v *Verbose) VisitBinary(w *walk.Walker, node *ast.BinaryExpression, metadata []walk.Metadata) {
	v.buffer.WriteString("(")
	w.Walk(node.Left, metadata)
	v.buffer.WriteString(" ")
	v.buffer.WriteString(node.Operator.String())
	v.buffer.WriteString(" ")
	w.Walk(node.Right, metadata)
	v.buffer.WriteString(")")
}

/*
func (v *Verbose3) VisitBlock(w *walk.Walker3, node *ast.BlockStatement, metadata []walk.Metadata) {
	for _, value := range node.List {
		w.Walk(value, metadata)
	}
}
*/

func (v *Verbose) VisitBoolean(w *walk.Walker, node *ast.BooleanLiteral, metadata []walk.Metadata) {
	v.buffer.WriteString(node.Literal)
}

func (v *Verbose) VisitCall(w *walk.Walker, node *ast.CallExpression, metadata []walk.Metadata) {
	w.Walk(node.Callee, metadata)
	v.buffer.WriteString("(")
	for _, value := range node.ArgumentList {
		w.Walk(value, metadata)
	}
	v.buffer.WriteString(")")
}

func (v *Verbose) VisitExpression(w *walk.Walker, node *ast.ExpressionStatement, metadata []walk.Metadata) {
	v.prependLevel()
	w.Walk(node.Expression, metadata)
	v.buffer.WriteString(";\n")
}

func (v *Verbose) VisitFor(w *walk.Walker, node *ast.ForStatement, metadata []walk.Metadata) {
	v.buffer.WriteString("for(")
	w.Walk(node.Initializer, metadata)
	v.buffer.WriteString(" ; ")
	w.Walk(node.Test, metadata)
	v.buffer.WriteString(" ; ")
	w.Walk(node.Update, metadata)
	v.buffer.WriteString(") {\n")
	v.level++
	w.Walk(node.Body, metadata)
	v.level--
	v.buffer.WriteString("}")
}

func (v *Verbose) VisitFunction(w *walk.Walker, node *ast.FunctionLiteral, metadata []walk.Metadata) {
	v.buffer.WriteString("function ")
	w.Walk(node.Name, metadata)
	v.buffer.WriteString("(")
	for _, value := range node.ParameterList.List {
		w.Walk(value, metadata)
	}
	v.buffer.WriteString(") {\n")
	v.level++
	w.Walk(node.Body, metadata)
	v.buffer.WriteString("}")

	v.level--

	for _, value := range node.DeclarationList {
		switch value := value.(type) {
		case *ast.FunctionDeclaration:
			w.Walk(value.Function, metadata)
		case *ast.VariableDeclaration:
			for _, value := range value.List {
				w.Walk(value, metadata)
			}
		default:
			panic(fmt.Errorf("Here be dragons: parseProgram.declaration(%T)", value))
		}
	}
}

func (v *Verbose) VisitIdentifier(w *walk.Walker, node *ast.Identifier, metadata []walk.Metadata) {
	v.buffer.WriteString(node.Name)
}

func (v *Verbose) VisitNumber(w *walk.Walker, node *ast.NumberLiteral, metadata []walk.Metadata) {
	v.buffer.WriteString(node.Literal)
}

/*
func (v *Verbose3) VisitSequence(w *walk.Walker3, node *ast.SequenceExpression, metadata []walk.Metadata) {
	println("[DEFAULT] Visiting sequence")
	for _, e := range node.Sequence {
		w.Walk(e, metadata)
		v.buffer.WriteString(";")
	}
}
*/

func (v *Verbose) VisitString(w *walk.Walker, node *ast.StringLiteral, metadata []walk.Metadata) {
	v.buffer.WriteString(node.Literal)
}

func (v *Verbose) VisitUnary(w *walk.Walker, node *ast.UnaryExpression, metadata []walk.Metadata) {
	if !node.Postfix {
		v.buffer.WriteString(node.Operator.String())
	}
	w.Walk(node.Operand, metadata)
	if node.Postfix {
		v.buffer.WriteString(node.Operator.String())
	}
}

func (v *Verbose) VisitWhile(w *walk.Walker, node *ast.WhileStatement, metadata []walk.Metadata) {
	v.buffer.WriteString("while(")
	w.Walk(node.Test, metadata)
	v.buffer.WriteString(") {\n")
	v.level++
	w.Walk(node.Body, metadata)
	v.level--
	v.buffer.WriteString("}")
}


func (v Verbose) ToString() string {
	return v.buffer.String()
}