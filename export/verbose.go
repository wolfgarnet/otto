package export

import (
	"bytes"

	"github.com/robertkrimen/otto/ast"
	"github.com/robertkrimen/otto/walk"
)

type Verbose struct {
	buffer bytes.Buffer
}

func (v *Verbose) Visit(node ast.Node) walk.Visitor {

	switch n := node.(type) {
	case *ast.BinaryExpression:
		v.buffer.WriteString(n.Operator.String())

	case *ast.Identifier:
		v.buffer.WriteString(n.Name)
	}

	return v
}

func (v *Verbose) ToString() string {
	return v.buffer.String()
}

type verboseBinary struct {

}