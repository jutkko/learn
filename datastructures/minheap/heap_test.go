package minheap

import "testing"

func TestPushAndPop(t *testing.T) {
	myMinHeap := &MinHeap{}

	myMinHeap.Push(9)
	myMinHeap.Push(3)
	myMinHeap.Push(-5)
	myMinHeap.Push(8)

	if myMinHeap.Pop() != -5 {
		t.Fatal("Push failed, expected -5")
	}

	if myMinHeap.Pop() != 3 {
		t.Fatal("Push failed, expected 3")
	}

	if myMinHeap.Pop() != 8 {
		t.Fatal("Push failed, expected 8")
	}
}

func TestSize(t *testing.T) {
	myMinHeap := &MinHeap{}

	myMinHeap.Push(9)
	myMinHeap.Push(3)
	myMinHeap.Push(-5)
	myMinHeap.Push(8)

	if myMinHeap.Size() != 4 {
		t.Fatal("Size is bad")
	}

	myMinHeap.Pop()

	if myMinHeap.Size() != 3 {
		t.Fatal("Size is bad")
	}
}
