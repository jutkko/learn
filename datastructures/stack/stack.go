package stack

import (
	"fmt"
	"log"
	"os"
)

type Node struct {
	child *Node
	Elem  int
}

func (n *Node) push(e int) *Node {
	if n == nil {
		return &Node{Elem: e}
	}

	return &Node{Elem: e, child: n}
}

func (n *Node) print() string {
	if n == nil {
		return ""
	}

	return fmt.Sprintf("%d ", n.Elem) + n.child.print()
}

func (n *Node) pop() *Node {
	if n == nil {
		return nil
	}

	return n.child
}

type Stack struct {
	head *Node
}

func (s *Stack) Push(e int) {
	s.head = s.head.push(e)
}

func (s *Stack) Pop() int {
	var result int

	if s.head != nil {
		result = s.head.Elem
	} else {
		log.Fatal("No elem left!")
		os.Exit(1)
	}

	s.head = s.head.pop()
	return result
}

func (s *Stack) Empty() bool {
	return s.head == nil
}

func (s *Stack) Print() string {
	return s.head.print()
}
