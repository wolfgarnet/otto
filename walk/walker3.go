package walk

import (
	"github.com/robertkrimen/otto/ast"
	"fmt"
)

type Walker3 struct {
	Visitor Visitor3
}

type Visitor3 interface {
	VisitArray(walker *Walker3, node *ast.ArrayLiteral, metadata []Metadata)
	VisitAssign(walker *Walker3, node *ast.AssignExpression, metadata []Metadata)
	VisitBad(walker *Walker3, node *ast.BadExpression, metadata []Metadata)
	VisitBadStatement(walker *Walker3, node *ast.BadStatement, metadata []Metadata)
	VisitBinary(walker *Walker3, node *ast.BinaryExpression, metadata []Metadata)
	VisitBlock(walker *Walker3, node *ast.BlockStatement, metadata []Metadata)
	VisitBoolean(walker *Walker3, node *ast.BooleanLiteral, metadata []Metadata)
	VisitBracket(walker *Walker3, node *ast.BracketExpression, metadata []Metadata)
	VisitBranch(walker *Walker3, node *ast.BranchStatement, metadata []Metadata)
	VisitCall(walker *Walker3, node *ast.CallExpression, metadata []Metadata)
	VisitCase(walker *Walker3, node *ast.CaseStatement, metadata []Metadata)
	VisitCatch(walker *Walker3, node *ast.CatchStatement, metadata []Metadata)
	VisitConditional(walker *Walker3, node *ast.ConditionalExpression, metadata []Metadata)
	VisitDebugger(walker *Walker3, node *ast.DebuggerStatement, metadata []Metadata)
	VisitDot(walker *Walker3, node *ast.DotExpression, metadata []Metadata)
	VisitDoWhile(walker *Walker3, node *ast.DoWhileStatement, metadata []Metadata)
	VisitEmpty(walker *Walker3, node *ast.EmptyStatement, metadata []Metadata)
	VisitExpression(walker *Walker3, node *ast.ExpressionStatement, metadata []Metadata)
	VisitForIn(walker *Walker3, node *ast.ForInStatement, metadata []Metadata)
	VisitFor(walker *Walker3, node *ast.ForStatement, metadata []Metadata)
	VisitFunction(walker *Walker3, node *ast.FunctionLiteral, metadata []Metadata)
	VisitIdentifier(walker *Walker3, node *ast.Identifier, metadata []Metadata)
	VisitIf(walker *Walker3, node *ast.IfStatement, metadata []Metadata)
	VisitLabelled(walker *Walker3, node *ast.LabelledStatement, metadata []Metadata)
	VisitNew(walker *Walker3, node *ast.NewExpression, metadata []Metadata)
	VisitNull(walker *Walker3, node *ast.NullLiteral, metadata []Metadata)
	VisitNumber(walker *Walker3, node *ast.NumberLiteral, metadata []Metadata)
	VisitObject(walker *Walker3, node *ast.ObjectLiteral, metadata []Metadata)
	VisitProgram(walker *Walker3, node *ast.Program, metadata []Metadata)
	VisitReturn(walker *Walker3, node *ast.ReturnStatement, metadata []Metadata)
	VisitRegex(walker *Walker3, node *ast.RegExpLiteral, metadata []Metadata)
	VisitSequence(walker *Walker3, node *ast.SequenceExpression, metadata []Metadata)
	VisitString(walker *Walker3, node *ast.StringLiteral, metadata []Metadata)
	VisitSwitch(walker *Walker3, node *ast.SwitchStatement, metadata []Metadata)
	VisitThis(walker *Walker3, node *ast.ThisExpression, metadata []Metadata)
	VisitThrow(walker *Walker3, node *ast.ThrowStatement, metadata []Metadata)
	VisitTry(walker *Walker3, node *ast.TryStatement, metadata []Metadata)
	VisitUnary(walker *Walker3, node *ast.UnaryExpression, metadata []Metadata)
	VisitVariable(walker *Walker3, node *ast.VariableExpression, metadata []Metadata)
	VisitVariableStatement(walker *Walker3, node *ast.VariableStatement, metadata []Metadata)
	VisitWhile(walker *Walker3, node *ast.WhileStatement, metadata []Metadata)
	VisitWith(walker *Walker3, node *ast.WithStatement, metadata []Metadata)
}

func (w *Walker3) Walk1(node ast.Node) {
	md := []Metadata{Metadata{nil}}
	w.Walk(node, md)
}

func (w *Walker3) Walk(node ast.Node, metadata []Metadata) {

	// Append the node
	metadata = append(metadata, Metadata{node})

	switch n := node.(type) {
	case *ast.ArrayLiteral:
		w.Visitor.VisitArray(w, n, metadata)
	case *ast.AssignExpression:
		w.Visitor.VisitAssign(w, n, metadata)
	case *ast.BadExpression:
		w.Visitor.VisitBad(w, n, metadata)
	case *ast.BadStatement:
		w.Visitor.VisitBadStatement(w, n, metadata)
	case *ast.BinaryExpression:
		w.Visitor.VisitBinary(w, n, metadata)
	case *ast.BlockStatement:
		w.Visitor.VisitBlock(w, n, metadata)
	case *ast.BooleanLiteral:
		w.Visitor.VisitBoolean(w, n, metadata)
	case *ast.BracketExpression:
		w.Visitor.VisitBracket(w, n, metadata)
	case *ast.BranchStatement:
		w.Visitor.VisitBranch(w, n, metadata)
	case *ast.CallExpression:
		w.Visitor.VisitCall(w, n, metadata)
	case *ast.CaseStatement:
		w.Visitor.VisitCase(w, n, metadata)
	case *ast.CatchStatement:
		w.Visitor.VisitCatch(w, n, metadata)
	case *ast.ConditionalExpression:
		w.Visitor.VisitConditional(w, n, metadata)
	case *ast.DebuggerStatement:
		w.Visitor.VisitDebugger(w, n, metadata)
	case *ast.DotExpression:
		w.Visitor.VisitDot(w, n, metadata)
	case *ast.DoWhileStatement:
		w.Visitor.VisitDoWhile(w, n, metadata)
	case *ast.EmptyStatement:
		w.Visitor.VisitEmpty(w, n, metadata)
	case *ast.ExpressionStatement:
		w.Visitor.VisitExpression(w, n, metadata)
	case *ast.ForInStatement:
		w.Visitor.VisitForIn(w, n, metadata)
	case *ast.ForStatement:
		w.Visitor.VisitFor(w, n, metadata)
	case *ast.FunctionLiteral:
		w.Visitor.VisitFunction(w, n, metadata)
	case *ast.Identifier:
		w.Visitor.VisitIdentifier(w, n, metadata)
	case *ast.IfStatement:
		w.Visitor.VisitIf(w, n, metadata)
	case *ast.LabelledStatement:
		w.Visitor.VisitLabelled(w, n, metadata)
	case *ast.NewExpression:
		w.Visitor.VisitNew(w, n, metadata)
	case *ast.NullLiteral:
		w.Visitor.VisitNull(w, n, metadata)
	case *ast.NumberLiteral:
		w.Visitor.VisitNumber(w, n, metadata)
	case *ast.ObjectLiteral:
		w.Visitor.VisitObject(w, n, metadata)
	case *ast.Program:
		w.Visitor.VisitProgram(w, n, metadata)
	case *ast.ReturnStatement:
		w.Visitor.VisitReturn(w, n, metadata)
	case *ast.RegExpLiteral:
		w.Visitor.VisitRegex(w, n, metadata)
	case *ast.SequenceExpression:
		w.Visitor.VisitSequence(w, n, metadata)
	case *ast.StringLiteral:
		w.Visitor.VisitString(w, n, metadata)
	case *ast.SwitchStatement:
		w.Visitor.VisitSwitch(w, n, metadata)
	case *ast.ThisExpression:
		w.Visitor.VisitThis(w, n, metadata)
	case *ast.ThrowStatement:
		w.Visitor.VisitThrow(w, n, metadata)
	case *ast.TryStatement:
		w.Visitor.VisitTry(w, n, metadata)
	case *ast.UnaryExpression:
		w.Visitor.VisitUnary(w, n, metadata)
	case *ast.VariableExpression:
		w.Visitor.VisitVariable(w, n, metadata)
	case *ast.VariableStatement:
		w.Visitor.VisitVariableStatement(w, n, metadata)
	case *ast.WhileStatement:
		w.Visitor.VisitWhile(w, n, metadata)
	case *ast.WithStatement:
		w.Visitor.VisitWith(w, n, metadata)
	}
}

type Visitor3Impl struct {

}

func (v *Visitor3Impl) VisitProgram(w *Walker3, node *ast.Program, metadata []Metadata) {
	fmt.Printf("[DEFAULT] Visiting program %v -- %v\n", node, displayMetadata(metadata))

	for _, e := range node.Body {
		w.Walk(e, metadata)
	}

	// Walking function and variable declarations
	for _, value := range node.DeclarationList {
		switch value := value.(type) {
		case *ast.FunctionDeclaration:
			w.Walk(value.Function, metadata)
		case *ast.VariableDeclaration:
			for _, value := range value.List {
				w.Walk(value, metadata)
			}
		default:
			panic(fmt.Errorf("Here be dragons: cmpl.parseProgram.DeclarationList(%T)", value))
		}
	}
}

func (v *Visitor3Impl) VisitArray(w *Walker3, node *ast.ArrayLiteral, metadata []Metadata) {
	fmt.Printf("[DEFAULT] Visiting array %v -- %v\n", node, displayMetadata(metadata))
}

func (v *Visitor3Impl) VisitAssign(w *Walker3, node *ast.AssignExpression, metadata []Metadata) {
	fmt.Printf("[DEFAULT] Visiting assign %v -- %v\n", node, displayMetadata(metadata))
	w.Walk(node.Left, metadata)
	w.Walk(node.Right, metadata)
}

func (v *Visitor3Impl) VisitBad(w *Walker3, node *ast.BadExpression, metadata []Metadata) {
	fmt.Printf("[DEFAULT] Visiting bad expression %v -- %v\n", node, displayMetadata(metadata))
}

func (v *Visitor3Impl) VisitBadStatement(w *Walker3, node *ast.BadStatement, metadata []Metadata) {
	fmt.Printf("[DEFAULT] Visiting bad statement %v -- %v\n", node, displayMetadata(metadata))
}

func (v *Visitor3Impl) VisitBinary(w *Walker3, node *ast.BinaryExpression, metadata []Metadata) {
	fmt.Printf("[DEFAULT] Visiting binary %v -- %v\n", node, displayMetadata(metadata))
	w.Walk(node.Left, metadata)
	w.Walk(node.Right, metadata)
}

func (v *Visitor3Impl) VisitBlock(w *Walker3, node *ast.BlockStatement, metadata []Metadata) {
	fmt.Printf("[DEFAULT] Visiting block %v -- %v\n", node, displayMetadata(metadata))
	for _, value := range node.List {
		w.Walk(value, metadata)
	}
}

func (v *Visitor3Impl) VisitBoolean(w *Walker3, node *ast.BooleanLiteral, metadata []Metadata) {
	fmt.Printf("[DEFAULT] Visiting boolean %v -- %v\n", node, displayMetadata(metadata))
}

func (v *Visitor3Impl) VisitBracket(w *Walker3, node *ast.BracketExpression, metadata []Metadata) {
	fmt.Printf("[DEFAULT] Visiting bracket %v -- %v\n", node, displayMetadata(metadata))
}

func (v *Visitor3Impl) VisitBranch(w *Walker3, node *ast.BranchStatement, metadata []Metadata) {
	fmt.Printf("[DEFAULT] Visiting branch %v -- %v\n", node, displayMetadata(metadata))
}

func (v *Visitor3Impl) VisitCall(w *Walker3, node *ast.CallExpression, metadata []Metadata) {
	fmt.Printf("[DEFAULT] Visiting call %v -- %v\n", node, displayMetadata(metadata))
	w.Walk(node.Callee, metadata)
	for _, value := range node.ArgumentList {
		w.Walk(value, metadata)
	}
}

func (v *Visitor3Impl) VisitCase(w *Walker3, node *ast.CaseStatement, metadata []Metadata) {
	fmt.Printf("[DEFAULT] Visiting case %v -- %v\n", node, displayMetadata(metadata))
}

func (v *Visitor3Impl) VisitCatch(w *Walker3, node *ast.CatchStatement, metadata []Metadata) {
	fmt.Printf("[DEFAULT] Visiting catch %v -- %v\n", node, displayMetadata(metadata))
}

func (v *Visitor3Impl) VisitConditional(w *Walker3, node *ast.ConditionalExpression, metadata []Metadata) {
	fmt.Printf("[DEFAULT] Visiting condition %v -- %v\n", node, displayMetadata(metadata))
}

func (v *Visitor3Impl) VisitDebugger(w *Walker3, node *ast.DebuggerStatement, metadata []Metadata) {
	fmt.Printf("[DEFAULT] Visiting  %v -- %v\n", node, displayMetadata(metadata))
}

func (v *Visitor3Impl) VisitDot(w *Walker3, node *ast.DotExpression, metadata []Metadata) {
	fmt.Printf("[DEFAULT] Visiting dot %v -- %v\n", node, displayMetadata(metadata))
}

func (v *Visitor3Impl) VisitDoWhile(w *Walker3, node *ast.DoWhileStatement, metadata []Metadata) {
	fmt.Printf("[DEFAULT] Visiting dowhile %v -- %v\n", node, displayMetadata(metadata))
}

func (v *Visitor3Impl) VisitEmpty(w *Walker3, node *ast.EmptyStatement, metadata []Metadata) {
	fmt.Printf("[DEFAULT] Visiting empty %v -- %v\n", node, displayMetadata(metadata))
}

func (v *Visitor3Impl) VisitExpression(w *Walker3, node *ast.ExpressionStatement, metadata []Metadata) {
	fmt.Printf("[DEFAULT] Visiting expression %v -- %v\n", node, displayMetadata(metadata))
	w.Walk(node.Expression, metadata)
}

func (v *Visitor3Impl) VisitForIn(w *Walker3, node *ast.ForInStatement, metadata []Metadata) {
	fmt.Printf("[DEFAULT] Visiting forin %v -- %v\n", node, displayMetadata(metadata))
}

func (v *Visitor3Impl) VisitFor(w *Walker3, node *ast.ForStatement, metadata []Metadata) {
	fmt.Printf("[DEFAULT] Visiting for %v -- %v\n", node, displayMetadata(metadata))
	w.Walk(node.Initializer, metadata)
	w.Walk(node.Test, metadata)
	w.Walk(node.Update, metadata)
	w.Walk(node.Body, metadata)
}

func (v *Visitor3Impl) VisitFunction(w *Walker3, node *ast.FunctionLiteral, metadata []Metadata) {
	fmt.Printf("[DEFAULT] Visiting function %v -- %v\n", node, displayMetadata(metadata))
	w.Walk(node.Name, metadata)
	for _, value := range node.ParameterList.List {
		w.Walk(value, metadata)
	}
	w.Walk(node.Body, metadata)

	for _, value := range node.DeclarationList {
		switch value := value.(type) {
		case *ast.FunctionDeclaration:
			w.Walk(value.Function, metadata)
		case *ast.VariableDeclaration:
			for _, value := range value.List {
				w.Walk(value, metadata)
			}
		default:
			panic(fmt.Errorf("Here be dragons: parseProgram.declaration(%T)", value))
		}
	}
}

func (v *Visitor3Impl) VisitIdentifier(w *Walker3, node *ast.Identifier, metadata []Metadata) {
	fmt.Printf("[DEFAULT] Visiting identifier %v -- %v\n", node, displayMetadata(metadata))
}

func (v *Visitor3Impl) VisitIf(w *Walker3, node *ast.IfStatement, metadata []Metadata) {
	fmt.Printf("[DEFAULT] Visiting if %v -- %v\n", node, displayMetadata(metadata))
}

func (v *Visitor3Impl) VisitLabelled(w *Walker3, node *ast.LabelledStatement, metadata []Metadata) {
	fmt.Printf("[DEFAULT] Visiting label %v -- %v\n", node, displayMetadata(metadata))
}

func (v *Visitor3Impl) VisitNew(w *Walker3, node *ast.NewExpression, metadata []Metadata) {
	fmt.Printf("[DEFAULT] Visiting new %v -- %v\n", node, displayMetadata(metadata))
}

func (v *Visitor3Impl) VisitNull(w *Walker3, node *ast.NullLiteral, metadata []Metadata) {
	fmt.Printf("[DEFAULT] Visiting null %v -- %v\n", node, displayMetadata(metadata))
}

func (v *Visitor3Impl) VisitNumber(w *Walker3, node *ast.NumberLiteral, metadata []Metadata) {
	fmt.Printf("[DEFAULT] Visiting number %v -- %v\n", node, displayMetadata(metadata))
}

func (v *Visitor3Impl) VisitObject(w *Walker3, node *ast.ObjectLiteral, metadata []Metadata) {
	fmt.Printf("[DEFAULT] Visiting object %v -- %v\n", node, displayMetadata(metadata))
}

func (v *Visitor3Impl) VisitReturn(w *Walker3, node *ast.ReturnStatement, metadata []Metadata) {
	fmt.Printf("[DEFAULT] Visiting return %v -- %v\n", node, displayMetadata(metadata))
}

func (v *Visitor3Impl) VisitRegex(w *Walker3, node *ast.RegExpLiteral, metadata []Metadata) {
	fmt.Printf("[DEFAULT] Visiting regex %v -- %v\n", node, displayMetadata(metadata))
}

func (v *Visitor3Impl) VisitSequence(w *Walker3, node *ast.SequenceExpression, metadata []Metadata) {
	fmt.Printf("[DEFAULT] Visiting sequence %v -- %v\n", node, displayMetadata(metadata))
	for _, e := range node.Sequence {
		w.Walk(e, metadata)
	}
}

func (v *Visitor3Impl) VisitString(w *Walker3, node *ast.StringLiteral, metadata []Metadata) {
	fmt.Printf("[DEFAULT] Visiting string %v -- %v\n", node, displayMetadata(metadata))
	// No op
}

func (v *Visitor3Impl) VisitSwitch(w *Walker3, node *ast.SwitchStatement, metadata []Metadata) {
	fmt.Printf("[DEFAULT] Visiting switch %v -- %v\n", node, displayMetadata(metadata))
}

func (v *Visitor3Impl) VisitThis(w *Walker3, node *ast.ThisExpression, metadata []Metadata) {
	fmt.Printf("[DEFAULT] Visiting this %v -- %v\n", node, displayMetadata(metadata))
}

func (v *Visitor3Impl) VisitThrow(w *Walker3, node *ast.ThrowStatement, metadata []Metadata) {
	fmt.Printf("[DEFAULT] Visiting throw %v -- %v\n", node, displayMetadata(metadata))
}

func (v *Visitor3Impl) VisitTry(w *Walker3, node *ast.TryStatement, metadata []Metadata) {
	fmt.Printf("[DEFAULT] Visiting try %v -- %v\n", node, displayMetadata(metadata))
}

func (v *Visitor3Impl) VisitUnary(w *Walker3, node *ast.UnaryExpression, metadata []Metadata) {
	fmt.Printf("[DEFAULT] Visiting unary %v -- %v\n", node, displayMetadata(metadata))
	w.Walk(node.Operand, metadata)
}

func (v *Visitor3Impl) VisitVariable(w *Walker3, node *ast.VariableExpression, metadata []Metadata) {
	fmt.Printf("[DEFAULT] Visiting variable expression %v -- %v\n", node, displayMetadata(metadata))
}

func (v *Visitor3Impl) VisitVariableStatement(w *Walker3, node *ast.VariableStatement, metadata []Metadata) {
	fmt.Printf("[DEFAULT] Visiting variable statement %v -- %v\n", node, displayMetadata(metadata))
}

func (v *Visitor3Impl) VisitWhile(w *Walker3, node *ast.WhileStatement, metadata []Metadata) {
	fmt.Printf("[DEFAULT] Visiting while %v -- %v\n", node, displayMetadata(metadata))
	w.Walk(node.Test, metadata)
	w.Walk(node.Body, metadata)
}

func (v *Visitor3Impl) VisitWith(w *Walker3, node *ast.WithStatement, metadata []Metadata) {
	fmt.Printf("[DEFAULT] Visiting with %v -- %v\n", node, displayMetadata(metadata))
}
