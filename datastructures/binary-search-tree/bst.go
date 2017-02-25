package bst

type Node struct {
	Left  *Node
	Right *Node
	Elem  int
}

func (n *Node) insert(newNode *Node) {
	if n.Elem == newNode.Elem {
		return
	}

	if n.Elem < newNode.Elem {
		if n.Right == nil {
			n.Right = newNode
			return
		} else {
			n.Right.insert(newNode)
		}
	} else {
		if n.Left == nil {
			n.Left = newNode
			return
		} else {
			n.Left.insert(newNode)
		}
	}
}

type Bst struct {
	Root *Node
}

func (b *Bst) Insert(value int) {
	if b.Root == nil {
		b.Root = &Node{Elem: value}
	} else {
		b.Root.insert(&Node{Elem: value})
	}

	return
}
