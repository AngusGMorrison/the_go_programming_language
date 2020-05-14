// Write a String method for the *tree type in gopl.io/ch4/treesort (§4.4) that reveals the
// sequence of values in the tree.
package main

import (
	"bytes"
	"fmt"
)

func main() {
	root := add(nil, 5)
	add(root, 3)
	add(root, 1)
	add(root, 7)
	fmt.Printf("%s\n", root.String())
}

type tree struct {
	value       int
	left, right *tree
}

func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		return &tree{value: value}
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

func (t *tree) String() string {
	buf := new(bytes.Buffer)
	stringify := func(t *tree) {
		if buf.Len() > 0 {
			fmt.Fprint(buf, " ")
		}
		fmt.Fprintf(buf, "%d", t.value)
	}
	eachNode(t, stringify)
	return buf.String()
}

func eachNode(t *tree, do func(t *tree)) {
	if t.left != nil {
		eachNode(t.left, do)
	}
	do(t)
	if t.right != nil {
		eachNode(t.right, do)
	}
}
