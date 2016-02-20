package ast

import (
	"fmt"
	"github.com/robertkrimen/otto/file"
)

// CommentPosition determines where the comment is in a given context
type CommentPosition int

const (
	_        CommentPosition = iota
	LEADING                  // Before the pertinent expression
	TRAILING                 // After the pertinent expression
	KEY                      // Before a key in an object
	COLON                    // After a colon in a field declaration
	FINAL                    // Final comments in a block, not belonging to a specific expression or the comment after a trailing , in an array or object literal
	IF                       // After an if keyword
	WHILE                    // After a while keyword
	DO                       // After do keyword
	FOR                      // After a for keyword
	WITH                     // After a with keyword
	TBD
)

// Comment contains the data of the comment
type Comment struct {
	Begin    file.Idx
	Text     string
	Position CommentPosition
}

// NewComment creates a new comment
func NewComment(text string, idx file.Idx) *Comment {
	comment := &Comment{
		Begin:    idx,
		Text:     text,
		Position: TBD,
	}

	return comment
}

// String returns a stringified version of the position
func (cp CommentPosition) String() string {
	switch cp {
	case LEADING:
		return "Leading"
	case TRAILING:
		return "Trailing"
	case KEY:
		return "Key"
	case COLON:
		return "Colon"
	case FINAL:
		return "Final"
	case IF:
		return "If"
	case WHILE:
		return "While"
	case DO:
		return "Do"
	case FOR:
		return "For"
	case WITH:
		return "With"
	default:
		return "???"
	}
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

func NewComments() *Comments {
	comments := &Comments{
		CommentMap: CommentMap{},
	}

	return comments
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
	fmt.Printf("RESETTIGN\n")
	c.wasLineBreak = false
}

// AddComment adds a comment to the view.
// Depending on the context, comments are added normally or as post line break.
func (c *Comments) AddComment(comment *Comment) {
	fmt.Printf("Adding comment, %v\n", c.wasLineBreak)
	if c.wasLineBreak {
		c.future = append(c.future, comment)
	} else {
		c.Comments = append(c.Comments, comment)
	}
}

func (c *Comments) MarkComments(position CommentPosition) {
	fmt.Printf("Marking comments as %v\n", position)
	for _, c := range c.Comments {
		if c.Position == TBD {
			c.Position = position
		}
	}
	for _, c := range c.future {
		if c.Position == TBD {
			c.Position = position
		}
	}
}

// Unset the current node and apply the scanned comments.
// Also marks the view as the next line.
func (c *Comments) Unset() {
	if c.Current != nil {
		c.applyComments(c.Current, TRAILING)
		c.Current = nil
	}
	c.wasLineBreak = false
}

// SetNode sets the current node of the view.
// It is skipped if the node is already set or if it is a part of the previous node.
// Scanned comments are linked to this node and future comments are promoted to normal comments.
// untilLineBreak marks the node as valid only until the next line break.
func (c *Comments) SetExpression(node Node, untilLineBreak bool) {
	// Skipping same node
	if c.Current == node {
		return
	}
	if c.Current != nil && c.Current.Idx1() == node.Idx1() {
		c.Current = node
		return
	}
	fmt.Printf("Current node is %v\n", node)

	c.Current = node
	c.UntilLineBreak = untilLineBreak

	// If a line break occurred, those regular comments must be linked to that node,
	// and any "future" comments must be marked as regular ones.
	c.applyComments(node, TRAILING)
	if c.wasLineBreak {
		fmt.Printf("Moving %v comments; %v\n", len(c.future), c.future)
		c.Comments = append(c.Comments, c.future...)
		c.future = nil
		c.wasLineBreak = false
	}
}

// Make sure the gathered comments are added
func (c *Comments) applyComments(node Node, position CommentPosition) {
	c.CommentMap.AddComments(node, c.Comments, position)
	c.Comments = nil
}

// AtLineBreak will mark subsequent comments as future.
// Also normal comments will be applied to the current node.
func (c *Comments) AtLineBreak() {
	if c.Current != nil {
		c.applyComments(c.Current, TRAILING)
	}

	// Promote future to leading comments
	/*if c.wasLineBreak {
		for _, c := range c.future {
			c.Position = LEADING
		}
		c.Comments = append(c.Comments, c.future...)
		c.future = nil
	}*/

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
func (cm CommentMap) AddComments(node Node, comments []*Comment, position CommentPosition) {
	for _, comment := range comments {
		if comment.Position == TBD {
			comment.Position = position
		}
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

// MoveComments moves comments with a given position from a node to another
func (cm CommentMap) MoveComments(from, to Node, position CommentPosition) {
	for i, c := range cm[from] {
		if c.Position == position {
			cm.AddComment(to, c)

			// Remove the comment from the "from" slice
			cm[from][i] = cm[from][len(cm[from])-1]
			cm[from][len(cm[from])-1] = nil
			cm[from] = cm[from][:len(cm[from])-1]
		}
	}
}
