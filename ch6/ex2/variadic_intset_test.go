package intset

import "testing"

const errorFmt = "Expected %v, got %v"

func createTestSet() *IntSet {
	set := IntSet{}
	set.AddAll(1, 2, 3)
	return &set
}

func TestLen(t *testing.T) {
	set := IntSet{}
	var len int
	if len = set.Len(); len != 0 {
		t.Errorf(errorFmt, 0, len)
	}
	set.Add(3)
	set.Add(2)
	if len = set.Len(); len != 2 {
		t.Errorf(errorFmt, 2, len)
	}
}

func TestRemove(t *testing.T) {
	set := createTestSet()
	set.Remove(2)
	if has := set.Has(2); has {
		t.Errorf(errorFmt, false, has)
	}
}

func TestClear(t *testing.T) {
	set := createTestSet()
	set.Clear()
	if len := set.Len(); len > 0 {
		t.Errorf(errorFmt, 0, len)
	}
}

func TestCopy(t *testing.T) {
	original := createTestSet()
	copy := original.Copy()
	if original == copy {
		t.Errorf("Original and copy addresses are equal")
	}
	if &original.words == &copy.words {
		t.Errorf("Original and copy use the same slice")
	}
}

func TestAddAll(t *testing.T) {
	set := createTestSet()
	set.AddAll(4, 5, 6)
	if !set.Has(4) || !set.Has(5) || !set.Has(6) {
		t.Errorf("AddAll failed to add all args to set")
	}
}
