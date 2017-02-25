package bst

import "fmt"

type Node struct {
	Left  *Node
	Right *Node
	Elem  int
}

func (n *Node) insert(newNode *Node) *Node {
	if n == nil {
		return newNode
	}

	if n.Elem < newNode.Elem {
		n.Right = n.Right.insert(newNode)
	} else {
		n.Left = n.Left.insert(newNode)
	}

	return n
}

func (n *Node) string() string {
	if n == nil {
		return ""
	}

	return n.Left.string() + fmt.Sprintf(" %d", n.Elem) + n.Right.string()
}

func (n *Node) min() *Node {
	if n == nil {
		return nil
	}

	if n.Left == nil {
		return n
	}

	return n.Left.min()
}

func (n *Node) exists(key int) bool {
	if n.Elem == key {
		return true
	}

	if n.Elem < key {
		return n.Right.exists(key)
	} else {
		return n.Left.exists(key)
	}
}

type Bst struct {
	Root *Node
}

func (b *Bst) Insert(value int) {
	b.Root = b.Root.insert(&Node{Elem: value})
}

func (b *Bst) String() string {
	return b.Root.string()
}

func (b *Bst) Min() *Node {
	return b.Root.min()
}

func (b *Bst) Exists(key int) bool {
	if b.Root == nil {
		return false
	}

	return b.Root.exists(key)
}
