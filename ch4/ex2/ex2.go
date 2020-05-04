/*
Write a program that prints the SHA256 hash of its standard input by default but supports a command
line flag to print the SHA384 or SHA512 hash intead.
*/

package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"log"
	"os"
)

var m = flag.Bool("m", false, "print SHA384 hash")
var l = flag.Bool("l", false, "print SHA512 hash")

func main() {
	const format = "%x\n"
	flag.Parse()

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text to hash: ")
	text, err := reader.ReadBytes('\n')
	if err != nil {
		log.Fatal(err)
	}

	if *m {
		fmt.Printf(format, sha512.Sum384(text))
	} else if *l {
		fmt.Printf(format, sha512.Sum512(text))
	} else {
		fmt.Printf(format, sha256.Sum256(text))
	}
}
