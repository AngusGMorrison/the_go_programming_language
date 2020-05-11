// Write a function expand(s string, f func(string) string) string that replaces each substring
// "$foo" within s by the text returned by f("foo").
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(expand("Here's a $foo sentence", expander))
}

func expand(s string, f func(string) string) string {
	// return strings.ReplaceAll(s, "$foo", f("foo"))
	const (
		substr    = "$foo"
		substrLen = len(substr)
	)

	var b strings.Builder
	for i := strings.Index(s, substr); i != -1; i = strings.Index(s, substr) {
		fmt.Println(s)
		b.WriteString(s[:i])
		b.WriteString(f("foo"))
		s = s[i+substrLen:]
	}
	b.WriteString(s)

	return b.String()
}

func expander(s string) string {
	return "EXPANDED"
}
