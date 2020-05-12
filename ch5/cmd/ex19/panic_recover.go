// Use panic and recover to write a function that contains no return statement yet returns a
// non-zero value.
package main

import "fmt"

func main() {
	fmt.Printf("%t\n", panicFunc())
}

func panicFunc() (hasReturnValue bool) {
	defer func() {
		if p := recover(); p != nil {
			hasReturnValue = true
		}
	}()
	panic("Panic!")
}
