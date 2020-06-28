// Rewrite topoSort to use maps instead of slices and eliminate the initial sort. Verify that the
// results, though nondeterministic, are valid topological orderings.
package main

import (
	"fmt"
)

// prereqs maps computer science courses to their prerequisites.
var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},

	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"data structures":       {"discrete maths"},
	"databases":             {"data structures"},
	"discrete maths":        {"intro to programming"},
	"formal languages":      {"discrete maths"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(item string)

	visitAll = func(item string) {
		if !seen[item] {
			seen[item] = true
			for _, prereq := range m[item] {
				visitAll(prereq)
			}
			order = append(order, item)
		}
	}

	for key := range m {
		visitAll(key)
	}
	return order
}
