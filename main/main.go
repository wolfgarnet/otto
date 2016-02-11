package main

import (
	"github.com/davecheney/profile"
	"github.com/robertkrimen/otto/parser"
)

func main() {
	defer profile.Start(profile.CPUProfile).Stop()

	parser.Proxy()
}
