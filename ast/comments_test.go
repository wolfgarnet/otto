package ast

import (
	"github.com/robertkrimen/otto/file"
	"testing"
)

func TestCommentMap(t *testing.T) {
	statement := &EmptyStatement{file.Idx(1)}
	comment := &Comment{1, "test"}

	cm := CommentMap{}
	cm.AddComment(statement, comment)

	if cm.Size() != 1 {
		t.Errorf("the number of comments is %v, not 1", cm.Size())
	}

	if len(cm[statement]) != 1 {
		t.Errorf("the number of comments is %v, not 1", cm.Size())
	}

	if cm[statement][0].Text != "test" {
		t.Errorf("the text is %v, not \"test\"", cm[statement][0].Text)
	}
}
