package walk

import (
	"github.com/robertkrimen/otto/ast"
	"bytes"
	"fmt"
	"reflect"
)

type Metadata struct {
	Parent ast.Node
}

func (md Metadata) String() string {
	return fmt.Sprintf("{parent:%v}", reflect.TypeOf(md.Parent))
}

func displayMetadata(metadata []Metadata) string {
	var buffer bytes.Buffer
	l := len(metadata)
	for i, md := range metadata {
		switch i {
		case l - 1:
			buffer.WriteString(fmt.Sprintf("[CURRENT] = %v", md))
		case l - 2:
			buffer.WriteString(fmt.Sprintf("[PARENT] = %v , ", md))
		default:
			buffer.WriteString(fmt.Sprintf("[%v] = %v , ", i, md))
		}

	}

	return buffer.String()
}