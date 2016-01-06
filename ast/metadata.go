package ast

import (
	"github.com/robertkrimen/otto/walk"
)

type Metadata struct {
	Parent Node
}

type MetadataMap map[Node][]*Metadata

func CreateMetadata(root Node) *MetadataMap {

}

type metadataWalker struct {
	walk.Walker3
	metadata MetadataMap
}


func (v *metadataWalker) VisitProgram(w *walk.Walker3, node *Program, parent Node) {
	for _, e := range node.Body {
		w.Walk(e, node)
	}

	// Walking function and variable declarations
	for _, value := range node.DeclarationList {
		switch value := value.(type) {
		case *FunctionDeclaration:
			w.Walk(value.Function, node)
		case *VariableDeclaration:
			for _, value := range value.List {
				w.Walk(value, node)
			}
		default:
			// No op
		}
	}
}

func (v *metadataWalker) VisitArray(w *walk.Walker3, node *ArrayLiteral, parent Node) {
	println("[DEFAULT] Visiting array", node)
}

func (v *metadataWalker) VisitAssign(w *walk.Walker3, node *AssignExpression, parent Node) {
	println("[DEFAULT] Visiting assign", node)
	w.Walk(node.Left, node)
	w.Walk(node.Right, node)
}

func (v *metadataWalker) VisitBad(w *walk.Walker3, node *BadExpression, parent Node) {
	println("[DEFAULT] Visiting bad expression", node)
}

func (v *metadataWalker) VisitBadStatement(w *walk.Walker3, node *BadStatement, parent Node) {
	println("[DEFAULT] Visiting bad statement", node)
}

func (v *metadataWalker) VisitBinary(w *walk.Walker3, node *BinaryExpression, parent Node) {
	println("[DEFAULT] Visiting binary", node)
	w.Walk(node.Left, node)
	w.Walk(node.Right, node)
}

func (v *metadataWalker) VisitBlock(w *walk.Walker3, node *BlockStatement, parent Node) {
	println("[DEFAULT] Visiting block", node)
	for _, value := range node.List {
		w.Walk(value, node)
	}
}

func (v *metadataWalker) VisitBoolean(w *walk.Walker3, node *BooleanLiteral, parent Node) {
	println("[DEFAULT] Visiting boolean", node)
}

func (v *metadataWalker) VisitBracket(w *walk.Walker3, node *BracketExpression, parent Node) {
	println("[DEFAULT] Visiting bracket", node)
}

func (v *metadataWalker) VisitBranch(w *walk.Walker3, node *BranchStatement, parent Node) {
	println("[DEFAULT] Visiting branch", node)
}

func (v *metadataWalker) VisitCall(w *walk.Walker3, node *CallExpression, parent Node) {
	println("[DEFAULT] Visiting call", node)
	w.Walk(node.Callee, node)
	for _, value := range node.ArgumentList {
		w.Walk(value, node)
	}
}

func (v *metadataWalker) VisitCase(w *walk.Walker3, node *CaseStatement, parent Node) {
	println("[DEFAULT] Visiting case", node)
}

func (v *metadataWalker) VisitCatch(w *walk.Walker3, node *CatchStatement, parent Node) {
	println("[DEFAULT] Visiting catch", node)
}

func (v *metadataWalker) VisitConditional(w *walk.Walker3, node *ConditionalExpression, parent Node) {
	println("[DEFAULT] Visiting condition", node)
}

func (v *metadataWalker) VisitDebugger(w *walk.Walker3, node *DebuggerStatement, parent Node) {
	println("[DEFAULT] Visiting ", node)
}

func (v *metadataWalker) VisitDot(w *walk.Walker3, node *DotExpression, parent Node) {
	println("[DEFAULT] Visiting dot", node)
}

func (v *metadataWalker) VisitDoWhile(w *walk.Walker3, node *DoWhileStatement, parent Node) {
	println("[DEFAULT] Visiting dowhile", node)
}

func (v *metadataWalker) VisitEmpty(w *walk.Walker3, node *EmptyStatement, parent Node) {
	println("[DEFAULT] Visiting empty", node)
}

func (v *metadataWalker) VisitExpression(w *walk.Walker3, node *ExpressionStatement, parent Node) {
	println("[DEFAULT] Visiting expression", node)
	w.Walk(node.Expression, node)
}

func (v *metadataWalker) VisitForIn(w *walk.Walker3, node *ForInStatement, parent Node) {
	println("[DEFAULT] Visiting forin", node)
}

func (v *metadataWalker) VisitFor(w *walk.Walker3, node *ForStatement, parent Node) {
	println("[DEFAULT] Visiting for", node)
	w.Walk(node.Initializer, node)
	w.Walk(node.Test, node)
	w.Walk(node.Update, node)
	w.Walk(node.Body, node)
}

func (v *metadataWalker) VisitFunction(w *walk.Walker3, node *FunctionLiteral, parent Node) {
	println("[DEFAULT] Visiting function", node)
	w.Walk(node.Name, node)
	for _, value := range node.ParameterList.List {
		w.Walk(value, node)
	}
	w.Walk(node.Body, node)

	for _, value := range node.DeclarationList {
		switch value := value.(type) {
		case *FunctionDeclaration:
			w.Walk(value.Function, node)
		case *VariableDeclaration:
			for _, value := range value.List {
				w.Walk(value, node)
			}
		default:
			// No op
		}
	}
}

func (v *metadataWalker) VisitIdentifier(w *walk.Walker3, node *Identifier, parent Node) {
	println("[DEFAULT] Visiting identifier", node)
}

func (v *metadataWalker) VisitIf(w *walk.Walker3, node *IfStatement, parent Node) {
	println("[DEFAULT] Visiting if", node)
}

func (v *metadataWalker) VisitLabelled(w *walk.Walker3, node *LabelledStatement, parent Node) {
	println("[DEFAULT] Visiting label", node)
}

func (v *metadataWalker) VisitNew(w *walk.Walker3, node *NewExpression, parent Node) {
	println("[DEFAULT] Visiting new", node)
}

func (v *metadataWalker) VisitNull(w *walk.Walker3, node *NullLiteral, parent Node) {
	println("[DEFAULT] Visiting null", node)
}

func (v *metadataWalker) VisitNumber(w *walk.Walker3, node *NumberLiteral, parent Node) {
	println("[DEFAULT] Visiting number", node)
}

func (v *metadataWalker) VisitObject(w *walk.Walker3, node *ObjectLiteral, parent Node) {
	println("[DEFAULT] Visiting object", node)
}

func (v *metadataWalker) VisitReturn(w *walk.Walker3, node *ReturnStatement, parent Node) {
	println("[DEFAULT] Visiting return", node)
}

func (v *metadataWalker) VisitRegex(w *walk.Walker3, node *RegExpLiteral, parent Node) {
	println("[DEFAULT] Visiting regex", node)
}

func (v *metadataWalker) VisitSequence(w *walk.Walker3, node *SequenceExpression, parent Node) {
	println("[DEFAULT] Visiting sequence", node)
	for _, e := range node.Sequence {
		w.Walk(e, node)
	}
}

func (v *metadataWalker) VisitString(w *walk.Walker3, node *StringLiteral, parent Node) {
	println("[DEFAULT] Visiting string", node)
	// No op
}

func (v *metadataWalker) VisitSwitch(w *walk.Walker3, node *SwitchStatement, parent Node) {
	println("[DEFAULT] Visiting switch", node)
}

func (v *metadataWalker) VisitThis(w *walk.Walker3, node *ThisExpression, parent Node) {
	println("[DEFAULT] Visiting this", node)
}

func (v *metadataWalker) VisitThrow(w *walk.Walker3, node *ThrowStatement, parent Node) {
	println("[DEFAULT] Visiting throw", node)
}

func (v *metadataWalker) VisitTry(w *walk.Walker3, node *TryStatement, parent Node) {
	println("[DEFAULT] Visiting try", node)
}

func (v *metadataWalker) VisitUnary(w *walk.Walker3, node *UnaryExpression, parent Node) {
	println("[DEFAULT] Visiting unary", node)
	w.Walk(node.Operand, node)
}

func (v *metadataWalker) VisitVariable(w *walk.Walker3, node *VariableExpression, parent Node) {
	println("[DEFAULT] Visiting variable expression", node)
}

func (v *metadataWalker) VisitVariableStatement(w *walk.Walker3, node *VariableStatement, parent Node) {
	println("[DEFAULT] Visiting variable statement", node)
}

func (v *metadataWalker) VisitWhile(w *walk.Walker3, node *WhileStatement, parent Node) {
	println("[DEFAULT] Visiting while", node)
	w.Walk(node.Test, node)
	w.Walk(node.Body, node)
}

func (v *metadataWalker) VisitWith(w *walk.Walker3, node *WithStatement, parent Node) {
	println("[DEFAULT] Visiting with", node)
}


