// Define a generic archive file-reading finction capable of reading ZIP files (archive/zip) and
// POSIX tar files (archive/tar). Use a registration mechanism similar to the one described above
// so that support for each file format can be plugged in using blank imports.
package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/angusgmorrison/the_go_programming_language/ch10/ex2/archive"
	_ "github.com/angusgmorrison/the_go_programming_language/ch10/ex2/archive/tar"
	_ "github.com/angusgmorrison/the_go_programming_language/ch10/ex2/archive/zip"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: %s path/to/archive\n", os.Args[0])
		os.Exit(1)
	}

	path := os.Args[1]
	fileInfo, err := archive.Decode(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Printf("%s contains:\n", filepath.Base(path))
	for _, file := range fileInfo {
		fmt.Println(file.Name())
		fmt.Printf("\t%d bytes\n", file.Size())
		fmt.Printf("\tDir? %t\n", file.IsDir())
	}
}
