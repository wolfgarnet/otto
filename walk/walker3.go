package walk

import (
	"github.com/robertkrimen/otto/ast"
	"fmt"
)

type Walker3 struct {
	Visitor Visitor3
}

type Visitor3 interface {
	VisitArray(walker *Walker3, node *ast.ArrayLiteral)
	VisitAssign(walker *Walker3, node *ast.AssignExpression)
	VisitBad(walker *Walker3, node *ast.BadExpression)
	VisitBadStatement(walker *Walker3, node *ast.BadStatement)
	VisitBinary(walker *Walker3, node *ast.BinaryExpression)
	VisitBlock(walker *Walker3, node *ast.BlockStatement)
	VisitBoolean(walker *Walker3, node *ast.BooleanLiteral)
	VisitBracket(walker *Walker3, node *ast.BracketExpression)
	VisitBranch(walker *Walker3, node *ast.BranchStatement)
	VisitCall(walker *Walker3, node *ast.CallExpression)
	VisitCase(walker *Walker3, node *ast.CaseStatement)
	VisitCatch(walker *Walker3, node *ast.CatchStatement)
	VisitConditional(walker *Walker3, node *ast.ConditionalExpression)
	VisitDebugger(walker *Walker3, node *ast.DebuggerStatement)
	VisitDot(walker *Walker3, node *ast.DotExpression)
	VisitDoWhile(walker *Walker3, node *ast.DoWhileStatement)
	VisitEmpty(walker *Walker3, node *ast.EmptyStatement)
	VisitExpression(walker *Walker3, node *ast.ExpressionStatement)
	VisitForIn(walker *Walker3, node *ast.ForInStatement)
	VisitFor(walker *Walker3, node *ast.ForStatement)
	VisitFunction(walker *Walker3, node *ast.FunctionLiteral)
	VisitIdentifier(walker *Walker3, node *ast.Identifier)
	VisitIf(walker *Walker3, node *ast.IfStatement)
	VisitLabelled(walker *Walker3, node *ast.LabelledStatement)
	VisitNew(walker *Walker3, node *ast.NewExpression)
	VisitNull(walker *Walker3, node *ast.NullLiteral)
	VisitNumber(walker *Walker3, node *ast.NumberLiteral)
	VisitObject(walker *Walker3, node *ast.ObjectLiteral)
	VisitProgram(walker *Walker3, node *ast.Program)
	VisitReturn(walker *Walker3, node *ast.ReturnStatement)
	VisitRegex(walker *Walker3, node *ast.RegExpLiteral)
	VisitSequence(walker *Walker3, node *ast.SequenceExpression)
	VisitString(walker *Walker3, node *ast.StringLiteral)
	VisitSwitch(walker *Walker3, node *ast.SwitchStatement)
	VisitThis(walker *Walker3, node *ast.ThisExpression)
	VisitThrow(walker *Walker3, node *ast.ThrowStatement)
	VisitTry(walker *Walker3, node *ast.TryStatement)
	VisitUnary(walker *Walker3, node *ast.UnaryExpression)
	VisitVariable(walker *Walker3, node *ast.VariableExpression)
	VisitVariableStatement(walker *Walker3, node *ast.VariableStatement)
	VisitWhile(walker *Walker3, node *ast.WhileStatement)
	VisitWith(walker *Walker3, node *ast.WithStatement)
}

func (w *Walker3) Walk(node ast.Node) {

	switch n := node.(type) {
	case *ast.ArrayLiteral:
		w.Visitor.VisitArray(w, n)
	case *ast.AssignExpression:
		w.Visitor.VisitAssign(w, n)
	case *ast.BadExpression:
		w.Visitor.VisitBad(w, n)
	case *ast.BadStatement:
		w.Visitor.VisitBadStatement(w, n)
	case *ast.BinaryExpression:
		w.Visitor.VisitBinary(w, n)
	case *ast.BlockStatement:
		w.Visitor.VisitBlock(w, n)
	case *ast.BooleanLiteral:
		w.Visitor.VisitBoolean(w, n)
	case *ast.BracketExpression:
		w.Visitor.VisitBracket(w, n)
	case *ast.BranchStatement:
		w.Visitor.VisitBranch(w, n)
	case *ast.CallExpression:
		w.Visitor.VisitCall(w, n)
	case *ast.CaseStatement:
		w.Visitor.VisitCase(w, n)
	case *ast.CatchStatement:
		w.Visitor.VisitCatch(w, n)
	case *ast.ConditionalExpression:
		w.Visitor.VisitConditional(w, n)
	case *ast.DebuggerStatement:
		w.Visitor.VisitDebugger(w, n)
	case *ast.DotExpression:
		w.Visitor.VisitDot(w, n)
	case *ast.DoWhileStatement:
		w.Visitor.VisitDoWhile(w, n)
	case *ast.EmptyStatement:
		w.Visitor.VisitEmpty(w, n)
	case *ast.ExpressionStatement:
		w.Visitor.VisitExpression(w, n)
	case *ast.ForInStatement:
		w.Visitor.VisitForIn(w, n)
	case *ast.ForStatement:
		w.Visitor.VisitFor(w, n)
	case *ast.FunctionLiteral:
		w.Visitor.VisitFunction(w, n)
	case *ast.Identifier:
		w.Visitor.VisitIdentifier(w, n)
	case *ast.IfStatement:
		w.Visitor.VisitIf(w, n)
	case *ast.LabelledStatement:
		w.Visitor.VisitLabelled(w, n)
	case *ast.NewExpression:
		w.Visitor.VisitNew(w, n)
	case *ast.NullLiteral:
		w.Visitor.VisitNull(w, n)
	case *ast.NumberLiteral:
		w.Visitor.VisitNumber(w, n)
	case *ast.ObjectLiteral:
		w.Visitor.VisitObject(w, n)
	case *ast.Program:
		w.Visitor.VisitProgram(w, n)
	case *ast.ReturnStatement:
		w.Visitor.VisitReturn(w, n)
	case *ast.RegExpLiteral:
		w.Visitor.VisitRegex(w, n)
	case *ast.SequenceExpression:
		w.Visitor.VisitSequence(w, n)
	case *ast.StringLiteral:
		w.Visitor.VisitString(w, n)
	case *ast.SwitchStatement:
		w.Visitor.VisitSwitch(w, n)
	case *ast.ThisExpression:
		w.Visitor.VisitThis(w, n)
	case *ast.ThrowStatement:
		w.Visitor.VisitThrow(w, n)
	case *ast.TryStatement:
		w.Visitor.VisitTry(w, n)
	case *ast.UnaryExpression:
		w.Visitor.VisitUnary(w, n)
	case *ast.VariableExpression:
		w.Visitor.VisitVariable(w, n)
	case *ast.VariableStatement:
		w.Visitor.VisitVariableStatement(w, n)
	case *ast.WhileStatement:
		w.Visitor.VisitWhile(w, n)
	case *ast.WithStatement:
		w.Visitor.VisitWith(w, n)
	}
}

type Visitor3Impl struct {

}

func (v *Visitor3Impl) VisitProgram(w *Walker3, node *ast.Program) {
	println("[DEFAULT] Visiting program", node)

	for _, e := range node.Body {
		w.Walk(e)
	}

	// Walking function and variable declarations
	for _, value := range node.DeclarationList {
		switch value := value.(type) {
		case *ast.FunctionDeclaration:
			w.Walk(value.Function)
		case *ast.VariableDeclaration:
			for _, value := range value.List {
				w.Walk(value)
			}
		default:
			panic(fmt.Errorf("Here be dragons: cmpl.parseProgram.DeclarationList(%T)", value))
		}
	}
}

func (v *Visitor3Impl) VisitArray(w *Walker3, node *ast.ArrayLiteral) {
	println("[DEFAULT] Visiting array", node)
}

func (v *Visitor3Impl) VisitAssign(w *Walker3, node *ast.AssignExpression) {
	println("[DEFAULT] Visiting assign", node)
}

func (v *Visitor3Impl) VisitBad(w *Walker3, node *ast.BadExpression) {
	println("[DEFAULT] Visiting bad expression", node)
}

func (v *Visitor3Impl) VisitBadStatement(w *Walker3, node *ast.BadStatement) {
	println("[DEFAULT] Visiting bad statement", node)
}

func (v *Visitor3Impl) VisitBinary(w *Walker3, node *ast.BinaryExpression) {
	println("[DEFAULT] Visiting binary", node)
	w.Walk(node.Left)
	w.Walk(node.Right)
}

func (v *Visitor3Impl) VisitBlock(w *Walker3, node *ast.BlockStatement) {
	println("[DEFAULT] Visiting block", node)
}

func (v *Visitor3Impl) VisitBoolean(w *Walker3, node *ast.BooleanLiteral) {
	println("[DEFAULT] Visiting boolean", node)
}

func (v *Visitor3Impl) VisitBracket(w *Walker3, node *ast.BracketExpression) {
	println("[DEFAULT] Visiting bracket", node)
}

func (v *Visitor3Impl) VisitBranch(w *Walker3, node *ast.BranchStatement) {
	println("[DEFAULT] Visiting branch", node)
}

func (v *Visitor3Impl) VisitCall(w *Walker3, node *ast.CallExpression) {
	println("[DEFAULT] Visiting call", node)
}

func (v *Visitor3Impl) VisitCase(w *Walker3, node *ast.CaseStatement) {
	println("[DEFAULT] Visiting case", node)
}

func (v *Visitor3Impl) VisitCatch(w *Walker3, node *ast.CatchStatement) {
	println("[DEFAULT] Visiting catch", node)
}

func (v *Visitor3Impl) VisitConditional(w *Walker3, node *ast.ConditionalExpression) {
	println("[DEFAULT] Visiting condition", node)
}

func (v *Visitor3Impl) VisitDebugger(w *Walker3, node *ast.DebuggerStatement) {
	println("[DEFAULT] Visiting ", node)
}

func (v *Visitor3Impl) VisitDot(w *Walker3, node *ast.DotExpression) {
	println("[DEFAULT] Visiting dot", node)
}

func (v *Visitor3Impl) VisitDoWhile(w *Walker3, node *ast.DoWhileStatement) {
	println("[DEFAULT] Visiting dowhile", node)
}

func (v *Visitor3Impl) VisitEmpty(w *Walker3, node *ast.EmptyStatement) {
	println("[DEFAULT] Visiting empty", node)
}

func (v *Visitor3Impl) VisitExpression(w *Walker3, node *ast.ExpressionStatement) {
	println("[DEFAULT] Visiting expression", node)
	w.Walk(node.Expression)
}

func (v *Visitor3Impl) VisitForIn(w *Walker3, node *ast.ForInStatement) {
	println("[DEFAULT] Visiting forin", node)
}

func (v *Visitor3Impl) VisitFor(w *Walker3, node *ast.ForStatement) {
	println("[DEFAULT] Visiting for", node)
}

func (v *Visitor3Impl) VisitFunction(w *Walker3, node *ast.FunctionLiteral) {
	println("[DEFAULT] Visiting function", node)
}

func (v *Visitor3Impl) VisitIdentifier(w *Walker3, node *ast.Identifier) {
	println("[DEFAULT] Visiting identifier", node)
}

func (v *Visitor3Impl) VisitIf(w *Walker3, node *ast.IfStatement) {
	println("[DEFAULT] Visiting if", node)
}

func (v *Visitor3Impl) VisitLabelled(w *Walker3, node *ast.LabelledStatement) {
	println("[DEFAULT] Visiting label", node)
}

func (v *Visitor3Impl) VisitNew(w *Walker3, node *ast.NewExpression) {
	println("[DEFAULT] Visiting new", node)
}

func (v *Visitor3Impl) VisitNull(w *Walker3, node *ast.NullLiteral) {
	println("[DEFAULT] Visiting null", node)
}

func (v *Visitor3Impl) VisitNumber(w *Walker3, node *ast.NumberLiteral) {
	println("[DEFAULT] Visiting number", node)
}

func (v *Visitor3Impl) VisitObject(w *Walker3, node *ast.ObjectLiteral) {
	println("[DEFAULT] Visiting object", node)
}

func (v *Visitor3Impl) VisitReturn(w *Walker3, node *ast.ReturnStatement) {
	println("[DEFAULT] Visiting return", node)
}

func (v *Visitor3Impl) VisitRegex(w *Walker3, node *ast.RegExpLiteral) {
	println("[DEFAULT] Visiting regex", node)
}

func (v *Visitor3Impl) VisitSequence(w *Walker3, node *ast.SequenceExpression) {
	println("[DEFAULT] Visiting sequence", node)
}

func (v *Visitor3Impl) VisitString(w *Walker3, node *ast.StringLiteral) {
	println("[DEFAULT] Visiting string", node)
}

func (v *Visitor3Impl) VisitSwitch(w *Walker3, node *ast.SwitchStatement) {
	println("[DEFAULT] Visiting switch", node)
}

func (v *Visitor3Impl) VisitThis(w *Walker3, node *ast.ThisExpression) {
	println("[DEFAULT] Visiting this", node)
}

func (v *Visitor3Impl) VisitThrow(w *Walker3, node *ast.ThrowStatement) {
	println("[DEFAULT] Visiting throw", node)
}

func (v *Visitor3Impl) VisitTry(w *Walker3, node *ast.TryStatement) {
	println("[DEFAULT] Visiting try", node)
}

func (v *Visitor3Impl) VisitUnary(w *Walker3, node *ast.UnaryExpression) {
	println("[DEFAULT] Visiting unary", node)
}

func (v *Visitor3Impl) VisitVariable(w *Walker3, node *ast.VariableExpression) {
	println("[DEFAULT] Visiting variable expression", node)
}

func (v *Visitor3Impl) VisitVariableStatement(w *Walker3, node *ast.VariableStatement) {
	println("[DEFAULT] Visiting variable statement", node)
}

func (v *Visitor3Impl) VisitWhile(w *Walker3, node *ast.WhileStatement) {
	println("[DEFAULT] Visiting while", node)
}

func (v *Visitor3Impl) VisitWith(w *Walker3, node *ast.WithStatement) {
	println("[DEFAULT] Visiting with", node)
}


