package export

import (
	"bytes"
	"fmt"
	"github.com/robertkrimen/otto/ast"
	"github.com/robertkrimen/otto/token"
	"github.com/robertkrimen/otto/walk"
	"strconv"
)

const (
	requiresBreak string = "requiresBreak"
)

type Verbose struct {
	walk.VisitorImpl
	buffer bytes.Buffer
	level  int
}

func (v *Verbose) write(str string) {
	v.buffer.WriteString(str)
}

func (v *Verbose) prependLevel() {
	for i := 0; i < v.level*3; i++ {
		v.buffer.WriteString(" ")
	}
}

func addBrackets(metadata []walk.Metadata, require bool) {
	md := walk.CurrentMetadata(metadata)
	md["brackets"] = require
}

func requireBrackets(metadata []walk.Metadata) bool {
	parent := walk.ParentMetadata(metadata)
	if parent == nil {
		return true
	}

	brackets, ok := parent["brackets"].(bool)
	if ok {
		return brackets
	}

	return true
}

func (v *Verbose) VisitArray(w *walk.Walker, node *ast.ArrayLiteral, metadata []walk.Metadata) walk.Metadata {
	v.buffer.WriteString("[")
	l := len(node.Value)
	for i, e := range node.Value {
		w.Walk(e, metadata)
		_, empty := e.(*ast.EmptyExpression)
		if i+1 < l || e == nil || empty {
			v.buffer.WriteString(", ")
		}
	}
	v.buffer.WriteString("]")

	return walk.CurrentMetadata(metadata)
}

func (v *Verbose) VisitAssign(w *walk.Walker, node *ast.AssignExpression, metadata []walk.Metadata) walk.Metadata {

	// Assignments in a binary expression must be grouped by parenthesis.
	_, group := walk.ParentMetadata(metadata).Parent().(*ast.BinaryExpression)

	if group {
		v.buffer.WriteString("(")
	}

	w.Walk(node.Left, metadata)
	v.buffer.WriteString(" ")
	v.buffer.WriteString(node.Operator.String())
	if node.Operator != token.ASSIGN {
		v.buffer.WriteString("=")
	}
	v.buffer.WriteString(" ")
	rightMD := w.Walk(node.Right, metadata)

	if group {
		v.buffer.WriteString(")")
	}

	md := walk.CurrentMetadata(metadata)
	md[requiresBreak] = rightMD[requiresBreak]
	return md
}

func (v *Verbose) VisitBinary(w *walk.Walker, node *ast.BinaryExpression, metadata []walk.Metadata) walk.Metadata {
	v.buffer.WriteString("(")
	w.Walk(node.Left, metadata)
	v.buffer.WriteString(" ")
	v.buffer.WriteString(node.Operator.String())
	v.buffer.WriteString(" ")
	w.Walk(node.Right, metadata)
	v.buffer.WriteString(")")

	return walk.CurrentMetadata(metadata)
}

func (v *Verbose) VisitBracket(w *walk.Walker, node *ast.BracketExpression, metadata []walk.Metadata) walk.Metadata {
	w.Walk(node.Left, metadata)
	v.buffer.WriteString("[")
	w.Walk(node.Member, metadata)
	v.buffer.WriteString("]")

	return walk.CurrentMetadata(metadata)
}

func (v *Verbose) VisitBranch(w *walk.Walker, node *ast.BranchStatement, metadata []walk.Metadata) walk.Metadata {
	v.prependLevel()
	v.buffer.WriteString(node.Token.String())
	if node.Label != nil {
		w.Walk(node.Label, metadata)
	}
	v.buffer.WriteString(";\n")

	return walk.CurrentMetadata(metadata)
}

func (v *Verbose) VisitBoolean(w *walk.Walker, node *ast.BooleanLiteral, metadata []walk.Metadata) walk.Metadata {
	v.buffer.WriteString(node.Literal)

	return walk.CurrentMetadata(metadata)
}

func (v *Verbose) VisitCall(w *walk.Walker, node *ast.CallExpression, metadata []walk.Metadata) walk.Metadata {
	calle, ok := node.Callee.(*ast.FunctionLiteral)
	if ok {
		v.buffer.WriteString("(")
	}
	w.Walk(node.Callee, metadata)
	v.buffer.WriteString("(")
	l := len(node.ArgumentList)
	for i, value := range node.ArgumentList {
		w.Walk(value, metadata)
		if i+1 < l {
			v.buffer.WriteString(", ")
		}
	}
	v.buffer.WriteString(")")

	if ok && calle != nil {
		v.buffer.WriteString(")")
	}

	return walk.CurrentMetadata(metadata)
}

func (v *Verbose) VisitDebugger(w *walk.Walker, node *ast.DebuggerStatement, metadata []walk.Metadata) walk.Metadata {
	v.buffer.WriteString("debugger")

	return walk.CurrentMetadata(metadata)
}

func (v *Verbose) VisitDot(w *walk.Walker, node *ast.DotExpression, metadata []walk.Metadata) walk.Metadata {
	w.Walk(node.Left, metadata)
	v.buffer.WriteString(".")
	w.Walk(node.Identifier, metadata)

	return walk.CurrentMetadata(metadata)
}

func (v *Verbose) VisitCase(w *walk.Walker, node *ast.CaseStatement, metadata []walk.Metadata) walk.Metadata {
	for _, e := range node.Consequent {
		w.Walk(e, metadata)
	}

	return walk.CurrentMetadata(metadata)
}

func (v *Verbose) VisitConditional(w *walk.Walker, node *ast.ConditionalExpression, metadata []walk.Metadata) walk.Metadata {
	v.buffer.WriteString("(")
	w.Walk(node.Test, metadata)
	v.buffer.WriteString(" ? ")
	w.Walk(node.Consequent, metadata)
	v.buffer.WriteString(" : ")
	w.Walk(node.Alternate, metadata)
	v.buffer.WriteString(")")

	return walk.CurrentMetadata(metadata)
}

func (v *Verbose) VisitDoWhile(w *walk.Walker, node *ast.DoWhileStatement, metadata []walk.Metadata) walk.Metadata {
	v.prependLevel()
	v.buffer.WriteString("do {\n")
	v.level++
	w.Walk(node.Body, metadata)
	v.level--

	v.prependLevel()
	v.buffer.WriteString("} while(")

	w.Walk(node.Test, metadata)

	v.buffer.WriteString(");\n")

	return walk.CurrentMetadata(metadata)
}

func (v *Verbose) VisitExpression(w *walk.Walker, node *ast.ExpressionStatement, metadata []walk.Metadata) walk.Metadata {
	v.prependLevel()
	md := w.Walk(node.Expression, metadata)
	rb, ok := md[requiresBreak].(bool)
	if !ok {
		rb = true
	}
	if rb {
		v.buffer.WriteString(";\n")
	}

	return walk.CurrentMetadata(metadata)
}

func (v *Verbose) VisitForIn(w *walk.Walker, node *ast.ForInStatement, metadata []walk.Metadata) walk.Metadata {
	v.prependLevel()
	v.buffer.WriteString("for(")
	w.Walk(node.Into, metadata)
	v.buffer.WriteString(" in ")
	w.Walk(node.Source, metadata)
	v.buffer.WriteString(") {\n")
	v.level++
	w.Walk(node.Body, metadata)
	v.level--

	v.prependLevel()
	v.buffer.WriteString("}\n")

	return walk.CurrentMetadata(metadata)
}

func (v *Verbose) VisitFor(w *walk.Walker, node *ast.ForStatement, metadata []walk.Metadata) walk.Metadata {
	v.prependLevel()
	v.buffer.WriteString("for(")

	w.Walk(node.Initializer, metadata)
	v.buffer.WriteString(" ; ")
	w.Walk(node.Test, metadata)
	v.buffer.WriteString(" ; ")
	w.Walk(node.Update, metadata)
	v.buffer.WriteString(") {\n")
	v.level++
	w.Walk(node.Body, metadata)
	v.level--

	v.prependLevel()
	v.buffer.WriteString("}\n")

	return walk.CurrentMetadata(metadata)
}

func (v *Verbose) VisitFunction(w *walk.Walker, node *ast.FunctionLiteral, metadata []walk.Metadata) walk.Metadata {

	if node.Name != nil {
		v.buffer.WriteString("function ")
		w.Walk(node.Name, metadata)
	} else {
		v.buffer.WriteString("function")
	}

	v.buffer.WriteString("(")
	l := len(node.ParameterList.List)
	for i, value := range node.ParameterList.List {
		w.Walk(value, metadata)
		if i+1 < l {
			v.buffer.WriteString(", ")
		}
	}
	v.buffer.WriteString(") {\n")
	v.level++
	w.Walk(node.Body, metadata)

	for _, value := range node.DeclarationList {
		switch value := value.(type) {
		case *ast.FunctionDeclaration:
			v.prependLevel()
			w.Walk(value.Function, metadata)
		case *ast.VariableDeclaration:
			/*for _, value := range value.List {
				w.Walk(value, metadata)
			}*/
		default:
			panic(fmt.Errorf("Here be dragons: parseProgram.declaration(%T)", value))
		}
	}

	v.level--
	v.prependLevel()
	v.buffer.WriteString("}")

	_, ok := metadata[len(metadata)-2]["parent"].(*ast.CallExpression)
	if !ok {
		v.buffer.WriteString("\n")
	}

	md := walk.CurrentMetadata(metadata)
	md[requiresBreak] = false
	return md
}

func (v *Verbose) VisitIdentifier(w *walk.Walker, node *ast.Identifier, metadata []walk.Metadata) walk.Metadata {
	v.buffer.WriteString(node.Name)

	return walk.CurrentMetadata(metadata)
}

func (v *Verbose) VisitIf(w *walk.Walker, node *ast.IfStatement, metadata []walk.Metadata) walk.Metadata {
	v.prependLevel()
	v.buffer.WriteString("if(")
	w.Walk(node.Test, metadata)
	v.buffer.WriteString(") {\n")
	v.level++
	w.Walk(node.Consequent, metadata)
	v.level--
	if node.Alternate != nil {
		v.prependLevel()
		v.buffer.WriteString("} else {\n")

		v.level++
		w.Walk(node.Alternate, metadata)
		v.level--
	}

	v.prependLevel()
	v.buffer.WriteString("}\n")

	return walk.CurrentMetadata(metadata)
}

func (v *Verbose) VisitNew(w *walk.Walker, node *ast.NewExpression, metadata []walk.Metadata) walk.Metadata {
	v.buffer.WriteString("new ")
	w.Walk(node.Callee, metadata)
	v.buffer.WriteString("(")
	l := len(node.ArgumentList)
	for i, e := range node.ArgumentList {
		w.Walk(e, metadata)

		if i+1 < l {
			v.buffer.WriteString(", ")
		}
	}
	v.buffer.WriteString(")")

	return walk.CurrentMetadata(metadata)
}

func (v *Verbose) VisitNull(w *walk.Walker, node *ast.NullLiteral, metadata []walk.Metadata) walk.Metadata {
	v.buffer.WriteString(node.Literal)

	return walk.CurrentMetadata(metadata)
}

func (v *Verbose) VisitNumber(w *walk.Walker, node *ast.NumberLiteral, metadata []walk.Metadata) walk.Metadata {
	v.buffer.WriteString(node.Literal)

	return walk.CurrentMetadata(metadata)
}

func (v *Verbose) VisitObject(w *walk.Walker, node *ast.ObjectLiteral, metadata []walk.Metadata) walk.Metadata {
	v.buffer.WriteString("{")
	l := len(node.Value)
	for i, value := range node.Value {
		v.buffer.WriteString(strconv.QuoteToASCII(value.Key))
		v.buffer.WriteString(": ")
		w.Walk(value.Value, metadata)
		if i+1 < l {
			v.buffer.WriteString(", ")
		}
	}
	v.buffer.WriteString("}")

	return walk.CurrentMetadata(metadata)
}

func (v *Verbose) VisitReturn(w *walk.Walker, node *ast.ReturnStatement, metadata []walk.Metadata) walk.Metadata {
	v.prependLevel()
	v.buffer.WriteString("return ")
	w.Walk(node.Argument, metadata)
	v.buffer.WriteString(";\n")

	return walk.CurrentMetadata(metadata)
}

func (v *Verbose) VisitRegex(w *walk.Walker, node *ast.RegExpLiteral, metadata []walk.Metadata) walk.Metadata {
	v.buffer.WriteString(node.Literal)
	v.buffer.WriteString(node.Flags)

	return walk.CurrentMetadata(metadata)
}

func (v *Verbose) VisitSequence(w *walk.Walker, node *ast.SequenceExpression, metadata []walk.Metadata) walk.Metadata {
	if len(node.Sequence) == 0 {
		return walk.CurrentMetadata(metadata)
	}

	_, addBrackets := walk.ParentMetadata(metadata).Parent().(*ast.AssignExpression)

	if addBrackets {
		v.buffer.WriteString("(")
	}

	l := len(node.Sequence)

	// Stupid special case
	if l > 0 {
		first, ok := node.Sequence[0].(*ast.VariableExpression)
		if ok && first != nil {
			idx := walk.FindVariable(metadata, first.Name)
			if idx == first.Idx {
				v.buffer.WriteString("var ")
			}
		}
	}

	for i, e := range node.Sequence {
		w.Walk(e, metadata)
		if i+1 < l {
			v.buffer.WriteString(", ")
		}
	}
	if addBrackets {
		v.buffer.WriteString(")")
	}

	return walk.CurrentMetadata(metadata)
}

func (v *Verbose) VisitString(w *walk.Walker, node *ast.StringLiteral, metadata []walk.Metadata) walk.Metadata {
	v.buffer.WriteString(node.Literal)

	return walk.CurrentMetadata(metadata)
}

func (v *Verbose) VisitSwitch(w *walk.Walker, node *ast.SwitchStatement, metadata []walk.Metadata) walk.Metadata {
	v.prependLevel()
	v.buffer.WriteString("switch(")
	w.Walk(node.Discriminant, metadata)
	v.buffer.WriteString(") {\n")
	for i, e := range node.Body {
		if i == node.Default {
			v.buffer.WriteString("default")
		} else {
			v.buffer.WriteString("case ")
			w.Walk(e.Test, metadata)
		}

		v.buffer.WriteString(":\n")

		v.level++
		w.Walk(e, metadata)
		v.level--
	}

	v.prependLevel()
	v.buffer.WriteString("}\n")

	return walk.CurrentMetadata(metadata)
}

func (v *Verbose) VisitThis(w *walk.Walker, node *ast.ThisExpression, metadata []walk.Metadata) walk.Metadata {
	v.buffer.WriteString("this")

	return walk.CurrentMetadata(metadata)
}

func (v *Verbose) VisitThrow(w *walk.Walker, node *ast.ThrowStatement, metadata []walk.Metadata) walk.Metadata {
	v.prependLevel()
	v.buffer.WriteString("throw ")
	w.Walk(node.Argument, metadata)
	v.buffer.WriteString(";\n")

	return walk.CurrentMetadata(metadata)
}

func (v *Verbose) VisitTry(w *walk.Walker, node *ast.TryStatement, metadata []walk.Metadata) walk.Metadata {
	v.prependLevel()
	v.buffer.WriteString("try {\n")

	v.level++
	w.Walk(node.Body, metadata)
	v.level--

	v.prependLevel()
	v.buffer.WriteString("}")

	if node.Catch != nil {
		v.prependLevel()
		v.buffer.WriteString(" catch(")

		w.Walk(node.Catch.Parameter, metadata)
		v.buffer.WriteString(") {\n")
		v.level++
		w.Walk(node.Catch.Body, metadata)
		v.level--

		v.prependLevel()
		v.buffer.WriteString("}")
	}

	if node.Finally != nil {
		v.prependLevel()
		v.buffer.WriteString(" finally {\n")

		v.level++
		w.Walk(node.Finally, metadata)
		v.level--

		v.prependLevel()
		v.buffer.WriteString("}")
	}

	v.buffer.WriteString("\n")

	return walk.CurrentMetadata(metadata)
}

func (v *Verbose) VisitUnary(w *walk.Walker, node *ast.UnaryExpression, metadata []walk.Metadata) walk.Metadata {
	if !node.Postfix {
		v.buffer.WriteString(node.Operator.String())
	}
	v.buffer.WriteString("(")
	w.Walk(node.Operand, metadata)
	v.buffer.WriteString(")")
	if node.Postfix {
		v.buffer.WriteString(node.Operator.String())
	}

	return walk.CurrentMetadata(metadata)
}

func (v *Verbose) VisitVariable(w *walk.Walker, node *ast.VariableExpression, metadata []walk.Metadata) walk.Metadata {
	v.buffer.WriteString(node.Name)
	if node.Initializer != nil {
		v.buffer.WriteString(" = ")
		rightMD := w.Walk(node.Initializer, metadata)

		md := walk.CurrentMetadata(metadata)
		md[requiresBreak] = rightMD[requiresBreak]
		return md
	}

	return walk.CurrentMetadata(metadata)
}

func (v *Verbose) VisitVariableStatement(w *walk.Walker, node *ast.VariableStatement, metadata []walk.Metadata) walk.Metadata {
	v.prependLevel()
	v.buffer.WriteString("var ")
	l := len(node.List)
	var last walk.Metadata
	for i, e := range node.List {
		last = w.Walk(e, metadata)
		if i+1 < l {
			v.buffer.WriteString(", ")
		}
	}

	rb := true
	if last != nil {
		b, ok := last[requiresBreak].(bool)
		if ok {
			rb = b
		}
	}

	if rb {
		v.buffer.WriteString(";\n")
	}

	return walk.CurrentMetadata(metadata)
}

func (v *Verbose) VisitWhile(w *walk.Walker, node *ast.WhileStatement, metadata []walk.Metadata) walk.Metadata {
	v.prependLevel()
	v.buffer.WriteString("while(")

	w.Walk(node.Test, metadata)
	v.buffer.WriteString(") {\n")
	v.level++
	w.Walk(node.Body, metadata)
	v.level--

	v.prependLevel()
	v.buffer.WriteString("}\n")

	return walk.CurrentMetadata(metadata)
}

func (v *Verbose) VisitWith(w *walk.Walker, node *ast.WithStatement, metadata []walk.Metadata) walk.Metadata {
	v.prependLevel()
	v.buffer.WriteString("with(")

	w.Walk(node.Object, metadata)
	v.buffer.WriteString(") {\n")
	v.level++
	w.Walk(node.Body, metadata)
	v.level--

	v.prependLevel()
	v.buffer.WriteString("}\n")

	return walk.CurrentMetadata(metadata)
}

func (v Verbose) ToString() string {
	return v.buffer.String()
}
