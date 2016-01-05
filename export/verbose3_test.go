package export

import (
	"testing"

	"github.com/robertkrimen/otto/walk"
	"github.com/robertkrimen/otto/parser"
)

func TestVerbose3(t *testing.T) {

	//src := `c = a + b - (1 + 2)`
	//src := `function fun(a) {alert("hej");c =1 + 2}`
	//src := `for(i = 0 ; i < 10 ; i++) {a = b + 2}`
	src := `while(true) {a = b + 2}`

	program, err := parser.ParseFile(nil, "", src, 0)
	if err != nil {
		t.Errorf("Failed, %v", err)
	}

	/*
	var x walk.Visitor3 = &Verbose3{}
	fmt.Printf("Type: %T\n", x)
	x.VisitIdentifier(nil, nil)
	*/

	verbose := &Verbose3{}
	walker := walk.Walker3{verbose}
	walker.Walk(program)
	println(verbose.ToString())
}

