package walk

import (
	"fmt"
	"github.com/robertkrimen/otto/ast"
)

// Walker can walk a given AST
type Walker struct {
	Visitor Visitor
}

// Visitor interface for the walker.
type Visitor interface {
	VisitArray(walker *Walker, node *ast.ArrayLiteral, metadata []Metadata)
	VisitAssign(walker *Walker, node *ast.AssignExpression, metadata []Metadata)
	VisitBad(walker *Walker, node *ast.BadExpression, metadata []Metadata)
	VisitBadStatement(walker *Walker, node *ast.BadStatement, metadata []Metadata)
	VisitBinary(walker *Walker, node *ast.BinaryExpression, metadata []Metadata)
	VisitBlock(walker *Walker, node *ast.BlockStatement, metadata []Metadata)
	VisitBoolean(walker *Walker, node *ast.BooleanLiteral, metadata []Metadata)
	VisitBracket(walker *Walker, node *ast.BracketExpression, metadata []Metadata)
	VisitBranch(walker *Walker, node *ast.BranchStatement, metadata []Metadata)
	VisitCall(walker *Walker, node *ast.CallExpression, metadata []Metadata)
	VisitCase(walker *Walker, node *ast.CaseStatement, metadata []Metadata)
	VisitCatch(walker *Walker, node *ast.CatchStatement, metadata []Metadata)
	VisitConditional(walker *Walker, node *ast.ConditionalExpression, metadata []Metadata)
	VisitDebugger(walker *Walker, node *ast.DebuggerStatement, metadata []Metadata)
	VisitDot(walker *Walker, node *ast.DotExpression, metadata []Metadata)
	VisitDoWhile(walker *Walker, node *ast.DoWhileStatement, metadata []Metadata)
	VisitEmpty(walker *Walker, node *ast.EmptyStatement, metadata []Metadata)
	VisitExpression(walker *Walker, node *ast.ExpressionStatement, metadata []Metadata)
	VisitForIn(walker *Walker, node *ast.ForInStatement, metadata []Metadata)
	VisitFor(walker *Walker, node *ast.ForStatement, metadata []Metadata)
	VisitFunction(walker *Walker, node *ast.FunctionLiteral, metadata []Metadata)
	VisitIdentifier(walker *Walker, node *ast.Identifier, metadata []Metadata)
	VisitIf(walker *Walker, node *ast.IfStatement, metadata []Metadata)
	VisitLabelled(walker *Walker, node *ast.LabelledStatement, metadata []Metadata)
	VisitNew(walker *Walker, node *ast.NewExpression, metadata []Metadata)
	VisitNull(walker *Walker, node *ast.NullLiteral, metadata []Metadata)
	VisitNumber(walker *Walker, node *ast.NumberLiteral, metadata []Metadata)
	VisitObject(walker *Walker, node *ast.ObjectLiteral, metadata []Metadata)
	VisitProgram(walker *Walker, node *ast.Program, metadata []Metadata)
	VisitReturn(walker *Walker, node *ast.ReturnStatement, metadata []Metadata)
	VisitRegex(walker *Walker, node *ast.RegExpLiteral, metadata []Metadata)
	VisitSequence(walker *Walker, node *ast.SequenceExpression, metadata []Metadata)
	VisitString(walker *Walker, node *ast.StringLiteral, metadata []Metadata)
	VisitSwitch(walker *Walker, node *ast.SwitchStatement, metadata []Metadata)
	VisitThis(walker *Walker, node *ast.ThisExpression, metadata []Metadata)
	VisitThrow(walker *Walker, node *ast.ThrowStatement, metadata []Metadata)
	VisitTry(walker *Walker, node *ast.TryStatement, metadata []Metadata)
	VisitUnary(walker *Walker, node *ast.UnaryExpression, metadata []Metadata)
	VisitVariable(walker *Walker, node *ast.VariableExpression, metadata []Metadata)
	VisitVariableStatement(walker *Walker, node *ast.VariableStatement, metadata []Metadata)
	VisitWhile(walker *Walker, node *ast.WhileStatement, metadata []Metadata)
	VisitWith(walker *Walker, node *ast.WithStatement, metadata []Metadata)
}

// Begin the walk of the given AST node
func (w *Walker) Begin(node ast.Node) {
	md := []Metadata{Metadata{nil}}
	w.Walk(node, md)
}

// Walk the AST, including metadata
func (w *Walker) Walk(node ast.Node, metadata []Metadata) {

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

// VisitorImpl is a default implementation of the Visitor interface
type VisitorImpl struct {
}

func (v *VisitorImpl) VisitProgram(w *Walker, node *ast.Program, metadata []Metadata) {
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

func (v *VisitorImpl) VisitArray(w *Walker, node *ast.ArrayLiteral, metadata []Metadata) {
	fmt.Printf("[DEFAULT] Visiting array %v -- %v\n", node, displayMetadata(metadata))
	for _, e := range node.Value {
		w.Walk(e, metadata)
	}
}

func (v *VisitorImpl) VisitAssign(w *Walker, node *ast.AssignExpression, metadata []Metadata) {
	fmt.Printf("[DEFAULT] Visiting assign %v -- %v\n", node, displayMetadata(metadata))
	w.Walk(node.Left, metadata)
	w.Walk(node.Right, metadata)
}

func (v *VisitorImpl) VisitBad(w *Walker, node *ast.BadExpression, metadata []Metadata) {
	fmt.Printf("[DEFAULT] Visiting bad expression %v -- %v\n", node, displayMetadata(metadata))
}

func (v *VisitorImpl) VisitBadStatement(w *Walker, node *ast.BadStatement, metadata []Metadata) {
	fmt.Printf("[DEFAULT] Visiting bad statement %v -- %v\n", node, displayMetadata(metadata))
}

func (v *VisitorImpl) VisitBinary(w *Walker, node *ast.BinaryExpression, metadata []Metadata) {
	fmt.Printf("[DEFAULT] Visiting binary %v -- %v\n", node, displayMetadata(metadata))
	w.Walk(node.Left, metadata)
	w.Walk(node.Right, metadata)
}

func (v *VisitorImpl) VisitBlock(w *Walker, node *ast.BlockStatement, metadata []Metadata) {
	fmt.Printf("[DEFAULT] Visiting block %v -- %v\n", node, displayMetadata(metadata))
	for _, value := range node.List {
		w.Walk(value, metadata)
	}
}

func (v *VisitorImpl) VisitBoolean(w *Walker, node *ast.BooleanLiteral, metadata []Metadata) {
	fmt.Printf("[DEFAULT] Visiting boolean %v -- %v\n", node, displayMetadata(metadata))
}

func (v *VisitorImpl) VisitBracket(w *Walker, node *ast.BracketExpression, metadata []Metadata) {
	fmt.Printf("[DEFAULT] Visiting bracket %v -- %v\n", node, displayMetadata(metadata))
	w.Walk(node.Left, metadata)
	w.Walk(node.Member, metadata)
}

func (v *VisitorImpl) VisitBranch(w *Walker, node *ast.BranchStatement, metadata []Metadata) {
	fmt.Printf("[DEFAULT] Visiting branch %v -- %v\n", node, displayMetadata(metadata))
	w.Walk(node.Label, metadata)
}

func (v *VisitorImpl) VisitCall(w *Walker, node *ast.CallExpression, metadata []Metadata) {
	fmt.Printf("[DEFAULT] Visiting call %v -- %v\n", node, displayMetadata(metadata))
	w.Walk(node.Callee, metadata)
	for _, value := range node.ArgumentList {
		w.Walk(value, metadata)
	}
}

func (v *VisitorImpl) VisitCase(w *Walker, node *ast.CaseStatement, metadata []Metadata) {
	fmt.Printf("[DEFAULT] Visiting case %v -- %v\n", node, displayMetadata(metadata))
	w.Walk(node.Test, metadata)
	for _, e := range node.Consequent {
		w.Walk(e, metadata)
	}
}

func (v *VisitorImpl) VisitCatch(w *Walker, node *ast.CatchStatement, metadata []Metadata) {
	fmt.Printf("[DEFAULT] Visiting catch %v -- %v\n", node, displayMetadata(metadata))
	w.Walk(node.Parameter, metadata)
	w.Walk(node.Body, metadata)
}

func (v *VisitorImpl) VisitConditional(w *Walker, node *ast.ConditionalExpression, metadata []Metadata) {
	fmt.Printf("[DEFAULT] Visiting condition %v -- %v\n", node, displayMetadata(metadata))
	w.Walk(node.Test, metadata)
	w.Walk(node.Consequent, metadata)
	w.Walk(node.Alternate, metadata)
}

func (v *VisitorImpl) VisitDebugger(w *Walker, node *ast.DebuggerStatement, metadata []Metadata) {
	fmt.Printf("[DEFAULT] Visiting  %v -- %v\n", node, displayMetadata(metadata))
}

func (v *VisitorImpl) VisitDot(w *Walker, node *ast.DotExpression, metadata []Metadata) {
	fmt.Printf("[DEFAULT] Visiting dot %v -- %v\n", node, displayMetadata(metadata))
	w.Walk(node.Left, metadata)
	w.Walk(node.Identifier, metadata)
}

func (v *VisitorImpl) VisitDoWhile(w *Walker, node *ast.DoWhileStatement, metadata []Metadata) {
	fmt.Printf("[DEFAULT] Visiting dowhile %v -- %v\n", node, displayMetadata(metadata))
	w.Walk(node.Test, metadata)
	w.Walk(node.Body, metadata)
}

func (v *VisitorImpl) VisitEmpty(w *Walker, node *ast.EmptyStatement, metadata []Metadata) {
	fmt.Printf("[DEFAULT] Visiting empty %v -- %v\n", node, displayMetadata(metadata))
}

func (v *VisitorImpl) VisitExpression(w *Walker, node *ast.ExpressionStatement, metadata []Metadata) {
	fmt.Printf("[DEFAULT] Visiting expression %v -- %v\n", node, displayMetadata(metadata))
	w.Walk(node.Expression, metadata)
}

func (v *VisitorImpl) VisitForIn(w *Walker, node *ast.ForInStatement, metadata []Metadata) {
	fmt.Printf("[DEFAULT] Visiting forin %v -- %v\n", node, displayMetadata(metadata))
	w.Walk(node.Into, metadata)
	w.Walk(node.Source, metadata)
	w.Walk(node.Body, metadata)
}

func (v *VisitorImpl) VisitFor(w *Walker, node *ast.ForStatement, metadata []Metadata) {
	fmt.Printf("[DEFAULT] Visiting for %v -- %v\n", node, displayMetadata(metadata))
	w.Walk(node.Initializer, metadata)
	w.Walk(node.Test, metadata)
	w.Walk(node.Update, metadata)
	w.Walk(node.Body, metadata)
}

func (v *VisitorImpl) VisitFunction(w *Walker, node *ast.FunctionLiteral, metadata []Metadata) {
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

func (v *VisitorImpl) VisitIdentifier(w *Walker, node *ast.Identifier, metadata []Metadata) {
	fmt.Printf("[DEFAULT] Visiting identifier %v -- %v\n", node, displayMetadata(metadata))
}

func (v *VisitorImpl) VisitIf(w *Walker, node *ast.IfStatement, metadata []Metadata) {
	fmt.Printf("[DEFAULT] Visiting if %v -- %v\n", node, displayMetadata(metadata))
	w.Walk(node.Test, metadata)
	w.Walk(node.Consequent, metadata)
	w.Walk(node.Alternate, metadata)
}

func (v *VisitorImpl) VisitLabelled(w *Walker, node *ast.LabelledStatement, metadata []Metadata) {
	fmt.Printf("[DEFAULT] Visiting label %v -- %v\n", node, displayMetadata(metadata))
	w.Walk(node.Label, metadata)
	w.Walk(node.Statement, metadata)
}

func (v *VisitorImpl) VisitNew(w *Walker, node *ast.NewExpression, metadata []Metadata) {
	fmt.Printf("[DEFAULT] Visiting new %v -- %v\n", node, displayMetadata(metadata))
	w.Walk(node.Callee, metadata)
	for _, e := range node.ArgumentList {
		w.Walk(e, metadata)
	}
}

func (v *VisitorImpl) VisitNull(w *Walker, node *ast.NullLiteral, metadata []Metadata) {
	fmt.Printf("[DEFAULT] Visiting null %v -- %v\n", node, displayMetadata(metadata))
}

func (v *VisitorImpl) VisitNumber(w *Walker, node *ast.NumberLiteral, metadata []Metadata) {
	fmt.Printf("[DEFAULT] Visiting number %v -- %v\n", node, displayMetadata(metadata))
}

func (v *VisitorImpl) VisitObject(w *Walker, node *ast.ObjectLiteral, metadata []Metadata) {
	fmt.Printf("[DEFAULT] Visiting object %v -- %v\n", node, displayMetadata(metadata))
}

func (v *VisitorImpl) VisitReturn(w *Walker, node *ast.ReturnStatement, metadata []Metadata) {
	fmt.Printf("[DEFAULT] Visiting return %v -- %v\n", node, displayMetadata(metadata))
	w.Walk(node.Argument, metadata)
}

func (v *VisitorImpl) VisitRegex(w *Walker, node *ast.RegExpLiteral, metadata []Metadata) {
	fmt.Printf("[DEFAULT] Visiting regex %v -- %v\n", node, displayMetadata(metadata))
}

func (v *VisitorImpl) VisitSequence(w *Walker, node *ast.SequenceExpression, metadata []Metadata) {
	fmt.Printf("[DEFAULT] Visiting sequence %v -- %v\n", node, displayMetadata(metadata))
	for _, e := range node.Sequence {
		w.Walk(e, metadata)
	}
}

func (v *VisitorImpl) VisitString(w *Walker, node *ast.StringLiteral, metadata []Metadata) {
	fmt.Printf("[DEFAULT] Visiting string %v -- %v\n", node, displayMetadata(metadata))
	// No op
}

func (v *VisitorImpl) VisitSwitch(w *Walker, node *ast.SwitchStatement, metadata []Metadata) {
	fmt.Printf("[DEFAULT] Visiting switch %v -- %v\n", node, displayMetadata(metadata))
	w.Walk(node.Discriminant, metadata)
	for _, e := range node.Body {
		w.Walk(e, metadata)
	}
}

func (v *VisitorImpl) VisitThis(w *Walker, node *ast.ThisExpression, metadata []Metadata) {
	fmt.Printf("[DEFAULT] Visiting this %v -- %v\n", node, displayMetadata(metadata))
}

func (v *VisitorImpl) VisitThrow(w *Walker, node *ast.ThrowStatement, metadata []Metadata) {
	fmt.Printf("[DEFAULT] Visiting throw %v -- %v\n", node, displayMetadata(metadata))
	w.Walk(node.Argument, metadata)
}

func (v *VisitorImpl) VisitTry(w *Walker, node *ast.TryStatement, metadata []Metadata) {
	fmt.Printf("[DEFAULT] Visiting try %v -- %v\n", node, displayMetadata(metadata))
	w.Walk(node.Body, metadata)
	w.Walk(node.Catch, metadata)
	w.Walk(node.Finally, metadata)
}

func (v *VisitorImpl) VisitUnary(w *Walker, node *ast.UnaryExpression, metadata []Metadata) {
	fmt.Printf("[DEFAULT] Visiting unary %v -- %v\n", node, displayMetadata(metadata))
	w.Walk(node.Operand, metadata)
}

func (v *VisitorImpl) VisitVariable(w *Walker, node *ast.VariableExpression, metadata []Metadata) {
	fmt.Printf("[DEFAULT] Visiting variable expression %v -- %v\n", node, displayMetadata(metadata))
	w.Walk(node.Initializer, metadata)
}

func (v *VisitorImpl) VisitVariableStatement(w *Walker, node *ast.VariableStatement, metadata []Metadata) {
	fmt.Printf("[DEFAULT] Visiting variable statement %v -- %v\n", node, displayMetadata(metadata))
	for _, e := range node.List {
		w.Walk(e, metadata)
	}
}

func (v *VisitorImpl) VisitWhile(w *Walker, node *ast.WhileStatement, metadata []Metadata) {
	fmt.Printf("[DEFAULT] Visiting while %v -- %v\n", node, displayMetadata(metadata))
	w.Walk(node.Test, metadata)
	w.Walk(node.Body, metadata)
}

func (v *VisitorImpl) VisitWith(w *Walker, node *ast.WithStatement, metadata []Metadata) {
	fmt.Printf("[DEFAULT] Visiting with %v -- %v\n", node, displayMetadata(metadata))
	w.Walk(node.Object, metadata)
	w.Walk(node.Body, metadata)
}
