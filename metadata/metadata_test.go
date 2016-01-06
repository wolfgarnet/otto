package ast

import (
	"testing"

	"github.com/robertkrimen/otto/parser"
)

func TestMetadata(t *testing.T) {
	src := `while(true) {a = b + 2}`

	program, err := parser.ParseFile(nil, "", src, 0)
	if err != nil {
		t.Errorf("Failed, %v", err)
	}

	md := CreateMetadata(program)
	md.Display()
}
