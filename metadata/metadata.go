package ast

import (
	"github.com/robertkrimen/otto/ast"
	"github.com/robertkrimen/otto/walk"
	"fmt"
	"reflect"
)

type Metadata struct {
	Parent ast.Node
}

type MetadataMap map[ast.Node]*Metadata

func (mdm MetadataMap) NewMetadata(node, parent ast.Node) *Metadata {
	md := &Metadata{parent}

	mdm[node] = md

	return md
}

func (mdm MetadataMap) Display() {
	fmt.Printf("Displaying metadata:\n")
	for k, v := range mdm {
		fmt.Printf(" * %v: [parent: %v]\n", reflect.TypeOf(k), reflect.TypeOf(v.Parent))
	}
	fmt.Printf("-------------------------\n")
}

func CreateMetadata(root ast.Node) MetadataMap {
	m := make(MetadataMap)
	mdvisitor := &metadataVisitor{
		metadata:m,
	}
	walker := walk.Walker3{mdvisitor}
	walker.Walk(root)

	return mdvisitor.metadata
}

type metadataVisitor struct {
	walk.Walker3
	metadata MetadataMap
}


func (v *metadataVisitor) VisitProgram(w *walk.Walker3, node *ast.Program) {
	for _, e := range node.Body {
		v.metadata.NewMetadata(e, node)
		w.Walk(e)
	}

	for _, value := range node.DeclarationList {
		switch value := value.(type) {
		case *ast.FunctionDeclaration:
			v.metadata.NewMetadata(value.Function, node)
			w.Walk(value.Function)
		case *ast.VariableDeclaration:
			for _, value := range value.List {
				v.metadata.NewMetadata(value, node)
				w.Walk(value)
			}
		default:
			// No op
		}
	}
}

func (v *metadataVisitor) VisitArray(w *walk.Walker3, node *ast.ArrayLiteral) {
	println("[DEFAULT] Visiting array", node)
}

func (v *metadataVisitor) VisitAssign(w *walk.Walker3, node *ast.AssignExpression) {
	println("[DEFAULT] Visiting assign", node)
	v.metadata.NewMetadata(node.Left, node)
	v.metadata.NewMetadata(node.Right, node)
	w.Walk(node.Left)
	w.Walk(node.Right)
}

func (v *metadataVisitor) VisitBad(w *walk.Walker3, node *ast.BadExpression) {
	println("[DEFAULT] Visiting bad expression", node)
}

func (v *metadataVisitor) VisitBadStatement(w *walk.Walker3, node *ast.BadStatement) {
	println("[DEFAULT] Visiting bad statement", node)
}

func (v *metadataVisitor) VisitBinary(w *walk.Walker3, node *ast.BinaryExpression) {
	println("[DEFAULT] Visiting binary", node)
	v.metadata.NewMetadata(node.Left, node)
	v.metadata.NewMetadata(node.Right, node)
	w.Walk(node.Left)
	w.Walk(node.Right)
}

func (v *metadataVisitor) VisitBlock(w *walk.Walker3, node *ast.BlockStatement) {
	println("[DEFAULT] Visiting block", node)
	for _, value := range node.List {
		w.Walk(value)
	}
}

func (v *metadataVisitor) VisitBoolean(w *walk.Walker3, node *ast.BooleanLiteral) {
	println("[DEFAULT] Visiting boolean", node)
}

func (v *metadataVisitor) VisitBracket(w *walk.Walker3, node *ast.BracketExpression) {
	println("[DEFAULT] Visiting bracket", node)
}

func (v *metadataVisitor) VisitBranch(w *walk.Walker3, node *ast.BranchStatement) {
	println("[DEFAULT] Visiting branch", node)
}

func (v *metadataVisitor) VisitCall(w *walk.Walker3, node *ast.CallExpression) {
	println("[DEFAULT] Visiting call", node)
	w.Walk(node.Callee)
	for _, value := range node.ArgumentList {
		w.Walk(value)
	}
}

func (v *metadataVisitor) VisitCase(w *walk.Walker3, node *ast.CaseStatement) {
	println("[DEFAULT] Visiting case", node)
}

func (v *metadataVisitor) VisitCatch(w *walk.Walker3, node *ast.CatchStatement) {
	println("[DEFAULT] Visiting catch", node)
}

func (v *metadataVisitor) VisitConditional(w *walk.Walker3, node *ast.ConditionalExpression) {
	println("[DEFAULT] Visiting condition", node)
}

func (v *metadataVisitor) VisitDebugger(w *walk.Walker3, node *ast.DebuggerStatement) {
	println("[DEFAULT] Visiting ", node)
}

func (v *metadataVisitor) VisitDot(w *walk.Walker3, node *ast.DotExpression) {
	println("[DEFAULT] Visiting dot", node)
}

func (v *metadataVisitor) VisitDoWhile(w *walk.Walker3, node *ast.DoWhileStatement) {
	println("[DEFAULT] Visiting dowhile", node)
}

func (v *metadataVisitor) VisitEmpty(w *walk.Walker3, node *ast.EmptyStatement) {
	println("[DEFAULT] Visiting empty", node)
}

func (v *metadataVisitor) VisitExpression(w *walk.Walker3, node *ast.ExpressionStatement) {
	println("[DEFAULT] Visiting expression", node)
	w.Walk(node.Expression)
}

func (v *metadataVisitor) VisitForIn(w *walk.Walker3, node *ast.ForInStatement) {
	println("[DEFAULT] Visiting forin", node)
}

func (v *metadataVisitor) VisitFor(w *walk.Walker3, node *ast.ForStatement) {
	println("[DEFAULT] Visiting for", node)
	w.Walk(node.Initializer)
	w.Walk(node.Test)
	w.Walk(node.Update)
	w.Walk(node.Body)
}

func (v *metadataVisitor) VisitFunction(w *walk.Walker3, node *ast.FunctionLiteral) {
	println("[DEFAULT] Visiting function", node)
	w.Walk(node.Name)
	for _, value := range node.ParameterList.List {
		w.Walk(value)
	}
	w.Walk(node.Body)

	for _, value := range node.DeclarationList {
		switch value := value.(type) {
		case *ast.FunctionDeclaration:
			w.Walk(value.Function)
		case *ast.VariableDeclaration:
			for _, value := range value.List {
				w.Walk(value)
			}
		default:
			// No op
		}
	}
}

func (v *metadataVisitor) VisitIdentifier(w *walk.Walker3, node *ast.Identifier) {
	println("[DEFAULT] Visiting identifier", node)
}

func (v *metadataVisitor) VisitIf(w *walk.Walker3, node *ast.IfStatement) {
	println("[DEFAULT] Visiting if", node)
}

func (v *metadataVisitor) VisitLabelled(w *walk.Walker3, node *ast.LabelledStatement) {
	println("[DEFAULT] Visiting label", node)
}

func (v *metadataVisitor) VisitNew(w *walk.Walker3, node *ast.NewExpression) {
	println("[DEFAULT] Visiting new", node)
}

func (v *metadataVisitor) VisitNull(w *walk.Walker3, node *ast.NullLiteral) {
	println("[DEFAULT] Visiting null", node)
}

func (v *metadataVisitor) VisitNumber(w *walk.Walker3, node *ast.NumberLiteral) {
	println("[DEFAULT] Visiting number", node)
}

func (v *metadataVisitor) VisitObject(w *walk.Walker3, node *ast.ObjectLiteral) {
	println("[DEFAULT] Visiting object", node)
}

func (v *metadataVisitor) VisitReturn(w *walk.Walker3, node *ast.ReturnStatement) {
	println("[DEFAULT] Visiting return", node)
}

func (v *metadataVisitor) VisitRegex(w *walk.Walker3, node *ast.RegExpLiteral) {
	println("[DEFAULT] Visiting regex", node)
}

func (v *metadataVisitor) VisitSequence(w *walk.Walker3, node *ast.SequenceExpression) {
	println("[DEFAULT] Visiting sequence", node)
	for _, e := range node.Sequence {
		w.Walk(e)
	}
}

func (v *metadataVisitor) VisitString(w *walk.Walker3, node *ast.StringLiteral) {
	println("[DEFAULT] Visiting string", node)
	// No op
}

func (v *metadataVisitor) VisitSwitch(w *walk.Walker3, node *ast.SwitchStatement) {
	println("[DEFAULT] Visiting switch", node)
}

func (v *metadataVisitor) VisitThis(w *walk.Walker3, node *ast.ThisExpression) {
	println("[DEFAULT] Visiting this", node)
}

func (v *metadataVisitor) VisitThrow(w *walk.Walker3, node *ast.ThrowStatement) {
	println("[DEFAULT] Visiting throw", node)
}

func (v *metadataVisitor) VisitTry(w *walk.Walker3, node *ast.TryStatement) {
	println("[DEFAULT] Visiting try", node)
}

func (v *metadataVisitor) VisitUnary(w *walk.Walker3, node *ast.UnaryExpression) {
	println("[DEFAULT] Visiting unary", node)
	w.Walk(node.Operand)
}

func (v *metadataVisitor) VisitVariable(w *walk.Walker3, node *ast.VariableExpression) {
	println("[DEFAULT] Visiting variable expression", node)
}

func (v *metadataVisitor) VisitVariableStatement(w *walk.Walker3, node *ast.VariableStatement) {
	println("[DEFAULT] Visiting variable statement", node)
}

func (v *metadataVisitor) VisitWhile(w *walk.Walker3, node *ast.WhileStatement) {
	println("[DEFAULT] Visiting while", node)
	w.Walk(node.Test)
	w.Walk(node.Body)
}

func (v *metadataVisitor) VisitWith(w *walk.Walker3, node *ast.WithStatement) {
	println("[DEFAULT] Visiting with", node)
}


