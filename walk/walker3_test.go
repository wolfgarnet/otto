package walk

import (
	"testing"
	"github.com/robertkrimen/otto/parser"
)

func TestWalker3(t *testing.T) {

	src := `function f() {a + b}`

	program, err := parser.ParseFile(nil, "", src, 0)
	if err != nil {
		t.Errorf("Failed, %v", err)
	}

	visitor := &Visitor3Impl{}
	walker := Walker3{visitor}
	walker.Walk1(program)
}
