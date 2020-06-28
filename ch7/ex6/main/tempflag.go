// Add support for Kelvin temperatures to tempflag.
package main

import (
	"flag"
	"fmt"

	"the_go_programming_language/ch7/ex6/tempconv"
)

var temp = tempconv.CelsiusFlag("temp", 20.0, "the temperature")

func main() {
	flag.Parse()
	fmt.Println(*temp)
}
