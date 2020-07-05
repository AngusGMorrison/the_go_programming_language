package intset

type IntSet interface {
	Has(x int) bool
	Add(x int)
	AddAll(ints ...int)
	UnionWith(t IntSet)
	String() string
	Ints() []int
}
