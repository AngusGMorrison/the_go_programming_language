// The instructor of the linear algebra course decides that calculus is now a prerequisite. Extend
// the topoSort function to report cycles.
package main

import (
	"fmt"
	"os"
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
	"linear algebra":        {"calculus"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func main() {
	courses, err := topoSort(prereqs)
	if err != nil {
		fmt.Fprintf(os.Stderr, "toposort: %v\n", err)
	}
	for i, course := range courses {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string][]string) ([]string, error) {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(item string) error
	// Track which items have already been seen in the current traversal with a recursion stack
	var recStack []string

	visitAll = func(item string) error {
		for _, seenItem := range recStack {
			if seenItem == item {
				return fmt.Errorf("cycle through %s: %v", item, append(recStack, item))
			}
		}
		recStack = append(recStack, item)

		if !seen[item] {
			seen[item] = true
			for _, prereq := range m[item] {
				if err := visitAll(prereq); err != nil {
					return err
				}
			}
			order = append(order, item)
			recStack = recStack[:len(recStack)-1]
		}
		return nil
	}

	for key := range m {
		if err := visitAll(key); err != nil {
			return nil, err
		}
	}
	return order, nil
}
