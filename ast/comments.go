package ast

import (
	"fmt"
	"github.com/robertkrimen/otto/file"
)

// Comment contains the data of the comment
type Comment struct {
	Begin file.Idx
	Text  string
}

// NewComment creates a new comment
func NewComment(text string, idx file.Idx) *Comment {
	comment := &Comment{
		Begin: idx,
		Text:  text,
	}

	return comment
}

// String returns a stringified version of the comment
func (c Comment) String() string {
	return fmt.Sprintf("Comment: %v", c.Text)
}

// Comments defines the current view of comments from the parser
type Comments struct {
	// Comments lists the comments scanned, not linked to a node yet
	Comments []*Comment
	// future lists the comments after a line break during a sequence of comments
	future []*Comment
	// Current is node for which comments are linked to
	Current Node
	// Maybe not necessary
	UntilLineBreak bool
	// wasLineBreak determines if a line break occured while scanning for comments
	wasLineBreak bool
	// CommentMap is a reference to the parser comment map
	CommentMap CommentMap
}

func (c *Comments) String() string {
	return fmt.Sprintf("NODE: %v, Comments: %v, Future: %v(LINEBREAK:%v)", c.Current, len(c.Comments), len(c.future), c.wasLineBreak)
}

// FetchAll returns all the currently scanned comments,
// including those from the next line
func (c *Comments) FetchAll() []*Comment {
	defer func() {
		c.Comments = nil
		c.future = nil
	}()

	return append(c.Comments, c.future...)
}

// Fetch returns all the currently scanned comments
func (c *Comments) Fetch() []*Comment {
	defer func() {
		c.Comments = nil
	}()

	return c.Comments
}

// ResetLineBreak marks the beginning of a new statement
func (c *Comments) ResetLineBreak() {
	c.wasLineBreak = false
}

// AddComment adds a comment to the view.
// Depending on the context, comments are added normally or as post line break.
func (c *Comments) AddComment(comment *Comment) {
	if c.wasLineBreak {
		c.future = append(c.future, comment)
	} else {
		c.Comments = append(c.Comments, comment)
	}
}

// Unset the current node and apply the scanned comments.
// Also marks the view as the next line.
func (c *Comments) Unset() {
	if c.Current != nil {
		c.applyComments(c.Current)
		c.Current = nil
	}
	c.wasLineBreak = false
}

// SetNode sets the current node of the view.
// It is skipped if the node is already set or if it is a part of the previous node.
// Scanned comments are linked to this node and future comments are promoted to normal comments.
// untilLineBreak marks the node as valid only until the next line break.
func (c *Comments) SetNode(node Node, untilLineBreak bool) {
	// Skipping same node
	if c.Current == node {
		return
	}
	if c.Current != nil && c.Current.Idx1() == node.Idx1() {
		c.Current = node
		return
	}

	c.Current = node
	c.UntilLineBreak = untilLineBreak

	// If a line break occurred, those regular comments must be linked to that node,
	// and any "future" comments must be marked as regular ones.
	c.applyComments(node)
	if c.wasLineBreak {
		c.Comments = append(c.Comments, c.future...)
		c.future = nil
		c.wasLineBreak = false
	}
}

// Make sure the gathered comments are added
func (c *Comments) applyComments(node Node) {
	c.CommentMap.AddComments(node, c.Comments)
	c.Comments = nil
}

// AtLineBreak will mark subsequent comments as future.
// Also normal comments will be applied to the current node.
func (c *Comments) AtLineBreak() {
	if c.Current != nil {
		c.applyComments(c.Current)
	}

	// Subsequent comments must not be associated with the current node
	if c.UntilLineBreak {
		c.Current = nil
	}

	c.wasLineBreak = true
}

// CommentMap is the data structure where all found comments are stored
type CommentMap map[Node][]*Comment

// AddComment adds a single comment to the map
func (cm CommentMap) AddComment(node Node, comment *Comment) {
	list := cm[node]
	list = append(list, comment)

	cm[node] = list
}

// AddComments adds a slice of comments, given a node and an updated position
func (cm CommentMap) AddComments(node Node, comments []*Comment) {
	for _, comment := range comments {
		cm.AddComment(node, comment)
	}
}

// Size returns the size of the map
func (cm CommentMap) Size() int {
	size := 0
	for _, comments := range cm {
		size += len(comments)
	}

	return size
}
