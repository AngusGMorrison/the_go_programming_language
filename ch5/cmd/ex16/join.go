// Write a variadic version of strings.Join.
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(join("O"))
	fmt.Println(join("O", "Hi", "Hey", "Howdy"))
}

func join(sep string, strs ...string) string {
	var b strings.Builder
	count := len(strs)
	for i, str := range strs {
		b.WriteString(str)
		if i < count-1 {
			b.WriteString(" ")
		}
	}
	return b.String()
}
