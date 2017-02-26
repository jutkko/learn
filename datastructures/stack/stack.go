package stack

import (
	"fmt"
	"log"
	"os"
)

type node struct {
	child *node
	elem  int
}

func (n *node) push(e int) *node {
	if n == nil {
		return &node{elem: e}
	}

	return &node{elem: e, child: n}
}

func (n *node) print() string {
	if n == nil {
		return ""
	}

	return fmt.Sprintf("%d ", n.elem) + n.child.print()
}

func (n *node) pop() *node {
	if n == nil {
		return nil
	}

	return n.child
}

type Stack struct {
	head *node
}

func (s *Stack) Push(e int) {
	s.head = s.head.push(e)
}

func (s *Stack) Pop() int {
	var result int

	if s.head != nil {
		result = s.head.elem
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
