package walk

import (
	"testing"
	"github.com/robertkrimen/otto/parser"
	"github.com/robertkrimen/otto/ast"
	"reflect"
)

type testWalker struct {
	VisitorImpl
	ancestors []Metadata
}

func (v *testWalker) VisitIdentifier(w *Walker, node *ast.Identifier, metadata []Metadata) {
	v.ancestors = metadata
}

func TestWalker(t *testing.T) {
	tests := []struct{
		src string
		size int
		parent reflect.Type
	}{
		{`1 + b`, 5, reflect.TypeOf((*ast.BinaryExpression)(nil))},
		{`c++`, 5, reflect.TypeOf((*ast.UnaryExpression)(nil))},
	}

	for i, test := range tests {
		program, err := parser.ParseFile(nil, "", test.src, 0)
		if err != nil {
			t.Errorf("[%v] Failed, %v", i, err)
		}

		visitor := &testWalker{}
		walker := Walker{visitor}
		walker.Begin(program)

		if test.size != len(visitor.ancestors) {
			t.Errorf("[%v] Failed, number of ancestors not correct, %v != %v", i, test.size, len(visitor.ancestors))
		}

		parent := visitor.ancestors[len(visitor.ancestors)-2].Parent
		typeOfParent := reflect.TypeOf(parent)

		if test.parent != typeOfParent {
			t.Errorf("[%v] Failed, parent not correct, %v != %v", i, test.parent, typeOfParent)
		}
	}
}
