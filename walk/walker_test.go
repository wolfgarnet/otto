package walk

import (
	"testing"
	"github.com/robertkrimen/otto/parser"
)

func TestWalker(t *testing.T) {

	src := `function f() {a + b}`

	program, err := parser.ParseFile(nil, "", src, 0)
	if err != nil {
		t.Errorf("Failed, %v", err)
	}

	visitor := &VisitorImpl{}
	walker := Walker{visitor}
	walker.Begin(program)
}
