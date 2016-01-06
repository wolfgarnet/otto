package walk

import (
	"github.com/robertkrimen/otto/ast"
	"fmt"
)

type Walker3 struct {
	Visitor Visitor3
}

type Visitor3 interface {
	VisitArray(walker *Walker3, node *ast.ArrayLiteral, parent ast.Node)
	VisitAssign(walker *Walker3, node *ast.AssignExpression, parent ast.Node)
	VisitBad(walker *Walker3, node *ast.BadExpression, parent ast.Node)
	VisitBadStatement(walker *Walker3, node *ast.BadStatement, parent ast.Node)
	VisitBinary(walker *Walker3, node *ast.BinaryExpression, parent ast.Node)
	VisitBlock(walker *Walker3, node *ast.BlockStatement, parent ast.Node)
	VisitBoolean(walker *Walker3, node *ast.BooleanLiteral, parent ast.Node)
	VisitBracket(walker *Walker3, node *ast.BracketExpression, parent ast.Node)
	VisitBranch(walker *Walker3, node *ast.BranchStatement, parent ast.Node)
	VisitCall(walker *Walker3, node *ast.CallExpression, parent ast.Node)
	VisitCase(walker *Walker3, node *ast.CaseStatement, parent ast.Node)
	VisitCatch(walker *Walker3, node *ast.CatchStatement, parent ast.Node)
	VisitConditional(walker *Walker3, node *ast.ConditionalExpression, parent ast.Node)
	VisitDebugger(walker *Walker3, node *ast.DebuggerStatement, parent ast.Node)
	VisitDot(walker *Walker3, node *ast.DotExpression, parent ast.Node)
	VisitDoWhile(walker *Walker3, node *ast.DoWhileStatement, parent ast.Node)
	VisitEmpty(walker *Walker3, node *ast.EmptyStatement, parent ast.Node)
	VisitExpression(walker *Walker3, node *ast.ExpressionStatement, parent ast.Node)
	VisitForIn(walker *Walker3, node *ast.ForInStatement, parent ast.Node)
	VisitFor(walker *Walker3, node *ast.ForStatement, parent ast.Node)
	VisitFunction(walker *Walker3, node *ast.FunctionLiteral, parent ast.Node)
	VisitIdentifier(walker *Walker3, node *ast.Identifier, parent ast.Node)
	VisitIf(walker *Walker3, node *ast.IfStatement, parent ast.Node)
	VisitLabelled(walker *Walker3, node *ast.LabelledStatement, parent ast.Node)
	VisitNew(walker *Walker3, node *ast.NewExpression, parent ast.Node)
	VisitNull(walker *Walker3, node *ast.NullLiteral, parent ast.Node)
	VisitNumber(walker *Walker3, node *ast.NumberLiteral, parent ast.Node)
	VisitObject(walker *Walker3, node *ast.ObjectLiteral, parent ast.Node)
	VisitProgram(walker *Walker3, node *ast.Program, parent ast.Node)
	VisitReturn(walker *Walker3, node *ast.ReturnStatement, parent ast.Node)
	VisitRegex(walker *Walker3, node *ast.RegExpLiteral, parent ast.Node)
	VisitSequence(walker *Walker3, node *ast.SequenceExpression, parent ast.Node)
	VisitString(walker *Walker3, node *ast.StringLiteral, parent ast.Node)
	VisitSwitch(walker *Walker3, node *ast.SwitchStatement, parent ast.Node)
	VisitThis(walker *Walker3, node *ast.ThisExpression, parent ast.Node)
	VisitThrow(walker *Walker3, node *ast.ThrowStatement, parent ast.Node)
	VisitTry(walker *Walker3, node *ast.TryStatement, parent ast.Node)
	VisitUnary(walker *Walker3, node *ast.UnaryExpression, parent ast.Node)
	VisitVariable(walker *Walker3, node *ast.VariableExpression, parent ast.Node)
	VisitVariableStatement(walker *Walker3, node *ast.VariableStatement, parent ast.Node)
	VisitWhile(walker *Walker3, node *ast.WhileStatement, parent ast.Node)
	VisitWith(walker *Walker3, node *ast.WithStatement, parent ast.Node)
}

func (w *Walker3) Walk(node ast.Node, parent ast.Node) {

	switch n := node.(type) {
	case *ast.ArrayLiteral:
		w.Visitor.VisitArray(w, n, parent)
	case *ast.AssignExpression:
		w.Visitor.VisitAssign(w, n, parent)
	case *ast.BadExpression:
		w.Visitor.VisitBad(w, n, parent)
	case *ast.BadStatement:
		w.Visitor.VisitBadStatement(w, n, parent)
	case *ast.BinaryExpression:
		w.Visitor.VisitBinary(w, n, parent)
	case *ast.BlockStatement:
		w.Visitor.VisitBlock(w, n, parent)
	case *ast.BooleanLiteral:
		w.Visitor.VisitBoolean(w, n, parent)
	case *ast.BracketExpression:
		w.Visitor.VisitBracket(w, n, parent)
	case *ast.BranchStatement:
		w.Visitor.VisitBranch(w, n, parent)
	case *ast.CallExpression:
		w.Visitor.VisitCall(w, n, parent)
	case *ast.CaseStatement:
		w.Visitor.VisitCase(w, n, parent)
	case *ast.CatchStatement:
		w.Visitor.VisitCatch(w, n, parent)
	case *ast.ConditionalExpression:
		w.Visitor.VisitConditional(w, n, parent)
	case *ast.DebuggerStatement:
		w.Visitor.VisitDebugger(w, n, parent)
	case *ast.DotExpression:
		w.Visitor.VisitDot(w, n, parent)
	case *ast.DoWhileStatement:
		w.Visitor.VisitDoWhile(w, n, parent)
	case *ast.EmptyStatement:
		w.Visitor.VisitEmpty(w, n, parent)
	case *ast.ExpressionStatement:
		w.Visitor.VisitExpression(w, n, parent)
	case *ast.ForInStatement:
		w.Visitor.VisitForIn(w, n, parent)
	case *ast.ForStatement:
		w.Visitor.VisitFor(w, n, parent)
	case *ast.FunctionLiteral:
		w.Visitor.VisitFunction(w, n, parent)
	case *ast.Identifier:
		w.Visitor.VisitIdentifier(w, n, parent)
	case *ast.IfStatement:
		w.Visitor.VisitIf(w, n, parent)
	case *ast.LabelledStatement:
		w.Visitor.VisitLabelled(w, n, parent)
	case *ast.NewExpression:
		w.Visitor.VisitNew(w, n, parent)
	case *ast.NullLiteral:
		w.Visitor.VisitNull(w, n, parent)
	case *ast.NumberLiteral:
		w.Visitor.VisitNumber(w, n, parent)
	case *ast.ObjectLiteral:
		w.Visitor.VisitObject(w, n, parent)
	case *ast.Program:
		w.Visitor.VisitProgram(w, n, parent)
	case *ast.ReturnStatement:
		w.Visitor.VisitReturn(w, n, parent)
	case *ast.RegExpLiteral:
		w.Visitor.VisitRegex(w, n, parent)
	case *ast.SequenceExpression:
		w.Visitor.VisitSequence(w, n, parent)
	case *ast.StringLiteral:
		w.Visitor.VisitString(w, n, parent)
	case *ast.SwitchStatement:
		w.Visitor.VisitSwitch(w, n, parent)
	case *ast.ThisExpression:
		w.Visitor.VisitThis(w, n, parent)
	case *ast.ThrowStatement:
		w.Visitor.VisitThrow(w, n, parent)
	case *ast.TryStatement:
		w.Visitor.VisitTry(w, n, parent)
	case *ast.UnaryExpression:
		w.Visitor.VisitUnary(w, n, parent)
	case *ast.VariableExpression:
		w.Visitor.VisitVariable(w, n, parent)
	case *ast.VariableStatement:
		w.Visitor.VisitVariableStatement(w, n, parent)
	case *ast.WhileStatement:
		w.Visitor.VisitWhile(w, n, parent)
	case *ast.WithStatement:
		w.Visitor.VisitWith(w, n, parent)
	}
}

type Visitor3Impl struct {

}

func (v *Visitor3Impl) VisitProgram(w *Walker3, node *ast.Program, parent ast.Node) {
	println("[DEFAULT] Visiting program", node)

	for _, e := range node.Body {
		w.Walk(e, node)
	}

	// Walking function and variable declarations
	for _, value := range node.DeclarationList {
		switch value := value.(type) {
		case *ast.FunctionDeclaration:
			w.Walk(value.Function, node)
		case *ast.VariableDeclaration:
			for _, value := range value.List {
				w.Walk(value, node)
			}
		default:
			panic(fmt.Errorf("Here be dragons: cmpl.parseProgram.DeclarationList(%T)", value))
		}
	}
}

func (v *Visitor3Impl) VisitArray(w *Walker3, node *ast.ArrayLiteral, parent ast.Node) {
	println("[DEFAULT] Visiting array", node)
}

func (v *Visitor3Impl) VisitAssign(w *Walker3, node *ast.AssignExpression, parent ast.Node) {
	println("[DEFAULT] Visiting assign", node)
	w.Walk(node.Left, node)
	w.Walk(node.Right, node)
}

func (v *Visitor3Impl) VisitBad(w *Walker3, node *ast.BadExpression, parent ast.Node) {
	println("[DEFAULT] Visiting bad expression", node)
}

func (v *Visitor3Impl) VisitBadStatement(w *Walker3, node *ast.BadStatement, parent ast.Node) {
	println("[DEFAULT] Visiting bad statement", node)
}

func (v *Visitor3Impl) VisitBinary(w *Walker3, node *ast.BinaryExpression, parent ast.Node) {
	println("[DEFAULT] Visiting binary", node)
	w.Walk(node.Left, node)
	w.Walk(node.Right, node)
}

func (v *Visitor3Impl) VisitBlock(w *Walker3, node *ast.BlockStatement, parent ast.Node) {
	println("[DEFAULT] Visiting block", node)
	for _, value := range node.List {
		w.Walk(value, node)
	}
}

func (v *Visitor3Impl) VisitBoolean(w *Walker3, node *ast.BooleanLiteral, parent ast.Node) {
	println("[DEFAULT] Visiting boolean", node)
}

func (v *Visitor3Impl) VisitBracket(w *Walker3, node *ast.BracketExpression, parent ast.Node) {
	println("[DEFAULT] Visiting bracket", node)
}

func (v *Visitor3Impl) VisitBranch(w *Walker3, node *ast.BranchStatement, parent ast.Node) {
	println("[DEFAULT] Visiting branch", node)
}

func (v *Visitor3Impl) VisitCall(w *Walker3, node *ast.CallExpression, parent ast.Node) {
	println("[DEFAULT] Visiting call", node)
	w.Walk(node.Callee, node)
	for _, value := range node.ArgumentList {
		w.Walk(value, node)
	}
}

func (v *Visitor3Impl) VisitCase(w *Walker3, node *ast.CaseStatement, parent ast.Node) {
	println("[DEFAULT] Visiting case", node)
}

func (v *Visitor3Impl) VisitCatch(w *Walker3, node *ast.CatchStatement, parent ast.Node) {
	println("[DEFAULT] Visiting catch", node)
}

func (v *Visitor3Impl) VisitConditional(w *Walker3, node *ast.ConditionalExpression, parent ast.Node) {
	println("[DEFAULT] Visiting condition", node)
}

func (v *Visitor3Impl) VisitDebugger(w *Walker3, node *ast.DebuggerStatement, parent ast.Node) {
	println("[DEFAULT] Visiting ", node)
}

func (v *Visitor3Impl) VisitDot(w *Walker3, node *ast.DotExpression, parent ast.Node) {
	println("[DEFAULT] Visiting dot", node)
}

func (v *Visitor3Impl) VisitDoWhile(w *Walker3, node *ast.DoWhileStatement, parent ast.Node) {
	println("[DEFAULT] Visiting dowhile", node)
}

func (v *Visitor3Impl) VisitEmpty(w *Walker3, node *ast.EmptyStatement, parent ast.Node) {
	println("[DEFAULT] Visiting empty", node)
}

func (v *Visitor3Impl) VisitExpression(w *Walker3, node *ast.ExpressionStatement, parent ast.Node) {
	println("[DEFAULT] Visiting expression", node)
	w.Walk(node.Expression, node)
}

func (v *Visitor3Impl) VisitForIn(w *Walker3, node *ast.ForInStatement, parent ast.Node) {
	println("[DEFAULT] Visiting forin", node)
}

func (v *Visitor3Impl) VisitFor(w *Walker3, node *ast.ForStatement, parent ast.Node) {
	println("[DEFAULT] Visiting for", node)
	w.Walk(node.Initializer, node)
	w.Walk(node.Test, node)
	w.Walk(node.Update, node)
	w.Walk(node.Body, node)
}

func (v *Visitor3Impl) VisitFunction(w *Walker3, node *ast.FunctionLiteral, parent ast.Node) {
	println("[DEFAULT] Visiting function", node)
	w.Walk(node.Name, node)
	for _, value := range node.ParameterList.List {
		w.Walk(value, node)
	}
	w.Walk(node.Body, node)

	for _, value := range node.DeclarationList {
		switch value := value.(type) {
		case *ast.FunctionDeclaration:
			w.Walk(value.Function, node)
		case *ast.VariableDeclaration:
			for _, value := range value.List {
				w.Walk(value, node)
			}
		default:
			panic(fmt.Errorf("Here be dragons: parseProgram.declaration(%T)", value))
		}
	}
}

func (v *Visitor3Impl) VisitIdentifier(w *Walker3, node *ast.Identifier, parent ast.Node) {
	println("[DEFAULT] Visiting identifier", node)
}

func (v *Visitor3Impl) VisitIf(w *Walker3, node *ast.IfStatement, parent ast.Node) {
	println("[DEFAULT] Visiting if", node)
}

func (v *Visitor3Impl) VisitLabelled(w *Walker3, node *ast.LabelledStatement, parent ast.Node) {
	println("[DEFAULT] Visiting label", node)
}

func (v *Visitor3Impl) VisitNew(w *Walker3, node *ast.NewExpression, parent ast.Node) {
	println("[DEFAULT] Visiting new", node)
}

func (v *Visitor3Impl) VisitNull(w *Walker3, node *ast.NullLiteral, parent ast.Node) {
	println("[DEFAULT] Visiting null", node)
}

func (v *Visitor3Impl) VisitNumber(w *Walker3, node *ast.NumberLiteral, parent ast.Node) {
	println("[DEFAULT] Visiting number", node)
}

func (v *Visitor3Impl) VisitObject(w *Walker3, node *ast.ObjectLiteral, parent ast.Node) {
	println("[DEFAULT] Visiting object", node)
}

func (v *Visitor3Impl) VisitReturn(w *Walker3, node *ast.ReturnStatement, parent ast.Node) {
	println("[DEFAULT] Visiting return", node)
}

func (v *Visitor3Impl) VisitRegex(w *Walker3, node *ast.RegExpLiteral, parent ast.Node) {
	println("[DEFAULT] Visiting regex", node)
}

func (v *Visitor3Impl) VisitSequence(w *Walker3, node *ast.SequenceExpression, parent ast.Node) {
	println("[DEFAULT] Visiting sequence", node)
	for _, e := range node.Sequence {
		w.Walk(e, node)
	}
}

func (v *Visitor3Impl) VisitString(w *Walker3, node *ast.StringLiteral, parent ast.Node) {
	println("[DEFAULT] Visiting string", node)
	// No op
}

func (v *Visitor3Impl) VisitSwitch(w *Walker3, node *ast.SwitchStatement, parent ast.Node) {
	println("[DEFAULT] Visiting switch", node)
}

func (v *Visitor3Impl) VisitThis(w *Walker3, node *ast.ThisExpression, parent ast.Node) {
	println("[DEFAULT] Visiting this", node)
}

func (v *Visitor3Impl) VisitThrow(w *Walker3, node *ast.ThrowStatement, parent ast.Node) {
	println("[DEFAULT] Visiting throw", node)
}

func (v *Visitor3Impl) VisitTry(w *Walker3, node *ast.TryStatement, parent ast.Node) {
	println("[DEFAULT] Visiting try", node)
}

func (v *Visitor3Impl) VisitUnary(w *Walker3, node *ast.UnaryExpression, parent ast.Node) {
	println("[DEFAULT] Visiting unary", node)
	w.Walk(node.Operand, node)
}

func (v *Visitor3Impl) VisitVariable(w *Walker3, node *ast.VariableExpression, parent ast.Node) {
	println("[DEFAULT] Visiting variable expression", node)
}

func (v *Visitor3Impl) VisitVariableStatement(w *Walker3, node *ast.VariableStatement, parent ast.Node) {
	println("[DEFAULT] Visiting variable statement", node)
}

func (v *Visitor3Impl) VisitWhile(w *Walker3, node *ast.WhileStatement, parent ast.Node) {
	println("[DEFAULT] Visiting while", node)
	w.Walk(node.Test, node)
	w.Walk(node.Body, node)
}

func (v *Visitor3Impl) VisitWith(w *Walker3, node *ast.WithStatement, parent ast.Node) {
	println("[DEFAULT] Visiting with", node)
}
