package parser

import "github.com/robertkrimen/otto/underscore"

func Proxy() {
	src := underscore.Source()

	for i := 0; i < 100; i++ {
		parser := newParser("", src)
		parser.mode = StoreComments
		parser.parse()
	}
}
