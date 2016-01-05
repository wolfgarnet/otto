package export

import (
	"github.com/robertkrimen/otto/walk"
	"bytes"
	"github.com/robertkrimen/otto/ast"
)

type Verbose2 struct {
	walk.Walker2

	buffer bytes.Buffer
}

func (v *Verbose2) VisitIdentifier(node *ast.Identifier) {
	println("HEJ1")
	v.buffer.WriteString(node.Name)
}

func (v Verbose2) ToString() string {
	return v.buffer.String()
}