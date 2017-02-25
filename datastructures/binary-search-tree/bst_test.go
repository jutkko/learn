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

func TestMinNoElem(t *testing.T) {
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

func TestExistsNoElem(t *testing.T) {
	bst := Bst{}

	if bst.Exists(0) {
		t.Fatalf("Exists is not right")
	}

	if bst.Exists(10) {
		t.Fatalf("Exists is not right")
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
		t.Fatalf("Exists is not right")
	}

	if bst.Exists(10) {
		t.Fatalf("Exists is not right")
	}
}

func TestSizeNoElem(t *testing.T) {
	bst := Bst{}

	if bst.Size() != 0 {
		t.Fatalf("Size is not right")
	}
}

func TestSize(t *testing.T) {
	bst := Bst{}
	bst.Insert(9)
	bst.Insert(0)
	bst.Insert(5)
	bst.Insert(5)
	bst.Insert(5)
	bst.Insert(-3)
	bst.Insert(1)
	bst.Insert(4)

	if bst.Size() != 8 {
		t.Fatalf("Size is not right")
	}
}

func TestHeightNoElem(t *testing.T) {
	bst := Bst{}

	if bst.Height() != 0 {
		t.Fatalf("Height is not right")
	}
}

func TestHeightElem(t *testing.T) {
	bst := Bst{}
	bst.Insert(9)
	bst.Insert(0)
	bst.Insert(5)
	bst.Insert(5)
	bst.Insert(5)
	bst.Insert(-3)
	bst.Insert(1)
	bst.Insert(4)

	if bst.Height() != 7 {
		t.Fatalf("Height is not right")
	}
}
