// Echo: prints its command line arguments

package main

import (
	"fmt"
	"os"
)

func main() {
	for i, arg := range os.Args {
		fmt.Printf("%d %s\n", i, arg)
	}
}
