package queue

import "testing"

func TestEnqueue(t *testing.T) {
	queue := &Queue{}

	queue.Enqueue(2)
	queue.Enqueue(5)

	if queue.Print() != "2 5 " {
		t.Fatal("Enqueue is bad")
	}
}

func TestDequeue(t *testing.T) {
	queue := &Queue{}

	queue.Enqueue(2)
	queue.Enqueue(5)
	queue.Enqueue(6)
	queue.Enqueue(1)

	elem := queue.Dequeue()
	if elem != 2 {
		t.Fatal("Dequeue is bad")
	}

	elem = queue.Dequeue()
	if elem != 5 {
		t.Fatal("Dequeue is bad")
	}

	elem = queue.Dequeue()
	if elem != 6 {
		t.Fatal("Dequeue is bad")
	}

	elem = queue.Dequeue()
	if elem != 1 {
		t.Fatal("Dequeue is bad")
	}
}

func TestEmpty(t *testing.T) {
	queue := &Queue{}

	if !queue.Empty() {
		t.Fatal("Empty is bad")
	}

	queue.Enqueue(6)
	queue.Enqueue(1)
	if queue.Empty() {
		t.Fatal("Empty is bad")
	}

	queue.Dequeue()
	queue.Dequeue()
	if !queue.Empty() {
		t.Fatal("Empty is bad")
	}
}
