package export

import (
	"bytes"
	"github.com/robertkrimen/otto/ast"
	"github.com/robertkrimen/otto/walk"
)

type Verbose3 struct {
	walk.Visitor3Impl
	buffer bytes.Buffer
}

func (v *Verbose3) VisitBinary(w *walk.Walker3, node *ast.BinaryExpression) {
	w.Walk(node.Left)
	v.buffer.WriteString(" ")
	v.buffer.WriteString(node.Operator.String())
	v.buffer.WriteString(" ")
	w.Walk(node.Right)
}

func (v *Verbose3) VisitIdentifier(w *walk.Walker3, node *ast.Identifier) {
	v.buffer.WriteString(node.Name)
}

func (v Verbose3) ToString() string {
	return v.buffer.String()
}