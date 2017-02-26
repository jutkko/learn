package queue

import (
	"fmt"
	"log"
	"os"
)

type node struct {
	prev *node
	next *node
	elem int
}

func (n *node) print() string {
	if n == nil {
		return ""
	}

	return fmt.Sprintf("%d ", n.elem) + n.next.print()
}

type Queue struct {
	head *node
	tail *node
}

func (q *Queue) Enqueue(e int) {
	// empty
	if q.head == nil {
		q.head = &node{elem: e}
		q.tail = q.head
		return
	}

	newNode := &node{elem: e}
	q.tail.next = newNode
	q.tail = newNode
}

func (q *Queue) Dequeue() int {
	if q.head == nil {
		log.Fatal("Queue empty")
		os.Exit(2)
	}

	result := q.head.elem

	// do this to prevent the tail to be left around
	if q.head == q.tail {
		q.head = nil
		q.tail = nil
	} else {
		q.head = q.head.next
	}

	return result
}

func (q *Queue) Empty() bool {
	return q.tail == nil
}

func (q *Queue) Print() string {
	return q.head.print()
}
