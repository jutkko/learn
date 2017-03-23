package lru

import (
	"testing"
)

func TestInsert(t *testing.T) {
	myLRU := New(2)

	myLRU.Insert(1, 2)
	myLRU.Insert(1, 3)
	value, err := myLRU.Get(1)
	if err != nil {
		t.Fatal("LRU get is not right")
	}

	if value != 3 {
		t.Fatal("LRU insert is not right")
	}

	if myLRU.Size() != 1 {
		t.Fatal("LRU size is not right")
	}

	myLRU.Insert(2, 3)
	myLRU.Insert(3, 3)

	if myLRU.Size() != 2 {
		t.Fatal("LRU size is not right")
	}
}

func TestZeroLimit(t *testing.T) {
	myLRU := New(0)

	myLRU.Insert(1, 2)
	myLRU.Insert(1, 3)
	_, err := myLRU.Get(1)

	if err != nil {
		if err.Error() != "No elem" {
			t.Fatal("LRU insert is not right")
		}
	}

	if myLRU.Size() != 0 {
		t.Fatal("LRU size is not right")
	}
}
