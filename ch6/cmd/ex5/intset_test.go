package intset

import "testing"

const errorFmt = "Expected %v, got %v"

func createTestSet() *IntSet {
	set := IntSet{}
	for i := 0; i < 3; i++ {
		set.Add(i + 1)
	}
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
