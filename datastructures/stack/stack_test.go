package stack

import "testing"

func TestPush(t *testing.T) {
	stack := Stack{}

	stack.Push(2)
	stack.Push(4)
	stack.Push(1)

	if stack.Print() != "1 4 2 " {
		t.Fatalf("Push is bad")
	}
}

func TestPop(t *testing.T) {
	stack := Stack{}

	stack.Push(2)
	stack.Push(4)
	stack.Push(1)

	elem := stack.Pop()
	if elem != 1 {
		t.Fatalf("Pop is bad, expecting 1 got %d", elem)
	}

	elem = stack.Pop()
	if elem != 4 {
		t.Fatalf("Pop is bad, expecting 1 got %d", elem)
	}

	elem = stack.Pop()
	if elem != 2 {
		t.Fatalf("Pop is bad, expecting 1 got %d", elem)
	}
}

func TestEmpty(t *testing.T) {
	stack := Stack{}

	if !stack.Empty() {
		t.Fatal("Empty is not working")
	}

	stack.Push(2)
	if stack.Empty() {
		t.Fatal("Empty is not working")
	}
}
