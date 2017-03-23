package lru

import "errors"

type Node struct {
	key   int
	value int
	prev  *Node
	next  *Node
}

type LRU struct {
	head  *Node
	tail  *Node
	size  int
	limit int
	table map[int]*Node
}

func (l *LRU) pop() {
	if l.head.next != l.tail {
		toRemove := l.head.next
		l.head.next = toRemove.next
		toRemove.next.prev = l.head

		delete(l.table, toRemove.key)
	}
}

func (l *LRU) reInsert(node *Node) {
	// remove node
	node.next.prev = node.prev
	node.prev.next = node.next

	// put to head
	node.next = l.head.next
	l.head.next = node

	node.next.prev = l.head.next
	node.prev = l.head
}

func New(limit int) *LRU {
	head := &Node{}
	tail := &Node{}

	head.next = tail
	tail.prev = head

	return &LRU{
		head:  head,
		tail:  tail,
		size:  0,
		limit: limit,
		table: make(map[int]*Node),
	}
}

func (l *LRU) Insert(k, v int) {
	if val, ok := l.table[k]; !ok {
		newNode := &Node{
			key:   k,
			value: v,
		}

		newNode.next = l.head.next
		l.head.next = newNode

		newNode.next.prev = l.head.next
		newNode.prev = l.head

		l.table[k] = newNode

		l.size += 1
		if l.size > l.limit {
			l.pop()
			l.size -= 1
		}
	} else {
		val.value = v
		l.reInsert(val)
	}
}

func (l *LRU) Get(k int) (int, error) {
	if val, ok := l.table[k]; ok {
		l.reInsert(val)

		return val.value, nil
	}

	return 0, errors.New("No elem")
}

func (l *LRU) Size() int {
	return l.size
}
