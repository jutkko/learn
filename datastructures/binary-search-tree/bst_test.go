package bst

import "testing"

func TestInsert(t *testing.T) {
	bst := Bst{}
	bst.Insert(5)
	bst.Insert(5)
	bst.Insert(1)
	bst.Insert(4)
	bst.Insert(9)

	node := &Node{}
	node.Right.insert(&Node{Elem: 4})
}
