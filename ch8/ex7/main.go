// Write a concurrent program that creates a local mirror of a web site, fetching each readable page
// and writing it to a directory on the local disk. Only pages within the original domain (for
// instance, golang.org) should be fetched. URLs within mirrored pages should be altered as needed
// so that they refer to the mirrored page, not the original.
package main

import (
	"flag"
	"fmt"
	"net/url"
	"os"
	"the_go_programming_language/ch8/ex7/mirror"
)

var rootDir string
var maxDownload uint64
var args []string

func init() {
	flag.StringVar(&rootDir, "root", "local", "mirror root directory")
	flag.Uint64Var(&maxDownload, "maxDownload", 10240, "max download size (bytes)")
	flag.Parse()
	args = flag.Args()
}

func main() {
	if len(args) != 1 {
		fmt.Println(args)
		fmt.Fprintf(os.Stderr, "usage: %s url\n", os.Args[0])
		os.Exit(1)
	}

	entrypoint, err := url.Parse(args[0])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	m := mirror.New(entrypoint, maxDownload, rootDir)
	if err = m.Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
