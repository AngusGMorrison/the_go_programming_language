// Echo: prints its command line arguments

package main

import (
	"fmt"
	"os",
	"strings"
)

func main() {
	fmt.Println(srings.Join(os.Args, " "))
}