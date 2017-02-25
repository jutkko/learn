package bst

import "testing"

func TestInsert(t *testing.T) {
	bst := Bst{}
	bst.Insert(9)
	bst.Insert(0)
	bst.Insert(5)
	bst.Insert(5)
	bst.Insert(5)
	bst.Insert(1)
	bst.Insert(4)

	if bst.String() != " 0 1 4 5 5 5 9" {
		t.Fatalf("Insert is not right")
	}
}

func TestMinNoMin(t *testing.T) {
	bst := Bst{}

	if bst.Min() != nil {
		t.Fatalf("Min is wrong")
	}
}

func TestMin(t *testing.T) {
	bst := Bst{}
	bst.Insert(9)
	bst.Insert(0)
	bst.Insert(5)
	bst.Insert(5)
	bst.Insert(5)
	bst.Insert(-3)
	bst.Insert(1)
	bst.Insert(4)

	if bst.Min().Elem != -3 {
		t.Fatalf("Min is not right: %d", bst.Min())
	}
}

func TestExists(t *testing.T) {
	bst := Bst{}
	bst.Insert(9)
	bst.Insert(0)
	bst.Insert(5)
	bst.Insert(5)
	bst.Insert(5)
	bst.Insert(-3)
	bst.Insert(1)
	bst.Insert(4)

	if !bst.Exists(0) {
		t.Fatalf("Find is not right")
	}
}
