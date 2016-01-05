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

func (v *Verbose3) VisitAssign(w *walk.Walker3, node *ast.AssignExpression) {
	w.Walk(node.Left)
	v.buffer.WriteString(" ")
	v.buffer.WriteString(node.Operator.String())
	v.buffer.WriteString(" ")
	w.Walk(node.Right)
}

func (v *Verbose3) VisitBinary(w *walk.Walker3, node *ast.BinaryExpression) {
	v.buffer.WriteString("(")
	w.Walk(node.Left)
	v.buffer.WriteString(" ")
	v.buffer.WriteString(node.Operator.String())
	v.buffer.WriteString(" ")
	w.Walk(node.Right)
	v.buffer.WriteString(")")
}

/*
func (v *Verbose3) VisitBlock(w *walk.Walker3, node *ast.BlockStatement) {
	for _, value := range node.List {
		w.Walk(value)
	}
}
*/

func (v *Verbose3) VisitBoolean(w *walk.Walker3, node *ast.BooleanLiteral) {
	v.buffer.WriteString(node.Literal)
}

func (v *Verbose3) VisitCall(w *walk.Walker3, node *ast.CallExpression) {
	w.Walk(node.Callee)
	v.buffer.WriteString("(")
	for _, value := range node.ArgumentList {
		w.Walk(value)
	}
	v.buffer.WriteString(")")
}

func (v *Verbose3) VisitExpression(w *walk.Walker3, node *ast.ExpressionStatement) {
	v.prependLevel()
	w.Walk(node.Expression)
	v.buffer.WriteString(";\n")
}

func (v *Verbose3) VisitFor(w *walk.Walker3, node *ast.ForStatement) {
	v.buffer.WriteString("for(")
	w.Walk(node.Initializer)
	v.buffer.WriteString(" ; ")
	w.Walk(node.Test)
	v.buffer.WriteString(" ; ")
	w.Walk(node.Update)
	v.buffer.WriteString(") {\n")
	v.level++
	w.Walk(node.Body)
	v.level--
	v.buffer.WriteString("}")
}

func (v *Verbose3) VisitFunction(w *walk.Walker3, node *ast.FunctionLiteral) {
	v.buffer.WriteString("function ")
	w.Walk(node.Name)
	v.buffer.WriteString("(")
	for _, value := range node.ParameterList.List {
		w.Walk(value)
	}
	v.buffer.WriteString(") {\n")
	v.level++
	w.Walk(node.Body)
	v.buffer.WriteString("}")

	v.level--

	for _, value := range node.DeclarationList {
		switch value := value.(type) {
		case *ast.FunctionDeclaration:
			w.Walk(value.Function)
		case *ast.VariableDeclaration:
			for _, value := range value.List {
				w.Walk(value)
			}
		default:
			panic(fmt.Errorf("Here be dragons: parseProgram.declaration(%T)", value))
		}
	}
}

func (v *Verbose3) VisitIdentifier(w *walk.Walker3, node *ast.Identifier) {
	v.buffer.WriteString(node.Name)
}

func (v *Verbose3) VisitNumber(w *walk.Walker3, node *ast.NumberLiteral) {
	v.buffer.WriteString(node.Literal)
}

/*
func (v *Verbose3) VisitSequence(w *walk.Walker3, node *ast.SequenceExpression) {
	println("[DEFAULT] Visiting sequence", node)
	for _, e := range node.Sequence {
		w.Walk(e)
		v.buffer.WriteString(";")
	}
}
*/

func (v *Verbose3) VisitString(w *walk.Walker3, node *ast.StringLiteral) {
	v.buffer.WriteString(node.Literal)
}

func (v *Verbose3) VisitUnary(w *walk.Walker3, node *ast.UnaryExpression) {
	if !node.Postfix {
		v.buffer.WriteString(node.Operator.String())
	}
	w.Walk(node.Operand)
	if node.Postfix {
		v.buffer.WriteString(node.Operator.String())
	}
}

func (v *Verbose3) VisitWhile(w *walk.Walker3, node *ast.WhileStatement) {
	v.buffer.WriteString("while(")
	w.Walk(node.Test)
	v.buffer.WriteString(") {\n")
	v.level++
	w.Walk(node.Body)
	v.level--
	v.buffer.WriteString("}")
}


func (v Verbose3) ToString() string {
	return v.buffer.String()
}