package export

import (
	"testing"
	"github.com/robertkrimen/otto/parser"
	"github.com/robertkrimen/otto/walk"
)

func TestVerbose(t *testing.T) {

	src := `a + b`

	program, err := parser.ParseFile(nil, "", src, 0)
	if err != nil {
		t.Errorf("Failed, %v", err)
	}

	verbose := &Verbose{}
	walker := walk.Walker{}
	walker.Walk(verbose, program)
	println("TEST", verbose.ToString())
}

