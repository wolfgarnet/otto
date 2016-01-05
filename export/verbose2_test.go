package export

import (
	"testing"
	"github.com/robertkrimen/otto/parser"
)

func TestVerbose2(t *testing.T) {

	src := `a + b`

	program, err := parser.ParseFile(nil, "", src, 0)
	if err != nil {
		t.Errorf("Failed, %v", err)
	}

	verbose := &Verbose2{}
	verbose.Walk(program)
	println("TEST", verbose.ToString())
}

