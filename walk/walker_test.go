package walk

import (
	"testing"
	"github.com/robertkrimen/otto/parser"
	"github.com/robertkrimen/otto/ast"
	"fmt"
)

type myVisitor struct {
}

func (mv myVisitor) Visit(node ast.Node) (w Visitor) {
	fmt.Printf("I AM VISITTING!\n")
	return myVisitor{}
}

func TestWalker(t *testing.T) {

	src := `a + b`

	program, err := parser.ParseFile(nil, "", src, 0)
	if err != nil {
		t.Errorf("Failed, %v", err)
	}

	walker := Walker{}
	walker.Walk(myVisitor{}, program)
}
