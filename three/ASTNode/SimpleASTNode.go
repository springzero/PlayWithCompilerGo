package ASTNode

import (
	"fmt"
	"strings"
)

type ASTNode struct {
	Typ         string
	Text        string
	children    []ASTNode
	parent      *ASTNode
	EpsilonNode *ASTNode
}

func (n *ASTNode) GetType() string {
	return n.Typ
}

func (n *ASTNode) GetText() string {
	return n.Text
}

func (n *ASTNode) Children() []ASTNode {
	return n.children
}

func (n *ASTNode) GetChildCount() int {
	return len(n.children)
}

func (n *ASTNode) GetChild(index int) ASTNode {
	return n.children[index]
}

func (n *ASTNode) GetChildren() []ASTNode {
	return n.children
}

func (n *ASTNode) Parent() *ASTNode {
	return n.parent
}

func (n *ASTNode) GetParent() *ASTNode {
	return n.parent
}

func (n *ASTNode) IsTerminal() bool {
	return len(n.children) == 0
}

func (n *ASTNode) AddChild(child ASTNode) {
	if child.isNamedNode() {
		n.children = append(n.children, child)
		child.parent = n
	} else {
		for _, node := range child.children {
			n.children = append(n.children, node)
			node.parent = n
		}
	}
}

func (n *ASTNode) setText(text string) {
	n.Text = text
}

func (n *ASTNode) isNamedNode() bool {
	if n.Typ != "" && len(n.Typ) > 1 && !strings.HasPrefix(n.Typ, "_") {
		return true
	}
	return false
}

func (n *ASTNode) dump(indent string) {
	var str string = indent + n.Typ
	if n.Text != "" {
		str += "(" + n.Text + ")"
	}
	fmt.Println(str)

	for _, child := range n.children {
		child.dump(indent + "\t")
	}
}
