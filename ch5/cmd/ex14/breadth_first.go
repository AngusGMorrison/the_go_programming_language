// Use the breadthFirst function to explore a different structure. For example, you could use the
// course dependencies from to topoSort example (a directed graph), the file system hierarchy on
// your computer (a tree), or a list of bus or subway routes downloaded from your city
// government's web site (an undirected graph).
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: %s /path/to/directory\n", os.Args[0])
		os.Exit(1)
	}
	breadthFirst(crawlDirectory, []string{os.Args[1]})
}

// crawlDirectory scans each file in a directory, prints its name, and returns the names of all
// directories.
func crawlDirectory(dirPath string) []string {
	// Check that dirPath is a directory
	fileInfo, err := os.Stat(dirPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return nil
	}
	if !fileInfo.IsDir() {
		fmt.Fprintf(os.Stderr, "%s is not a directory\n", dirPath)
		return nil
	}

	// Iterate over files in directory, appending any new directories to the slice of directory
	// names to return.
	dirFiles, err := ioutil.ReadDir(dirPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return nil
	}

	var foundDirs []string
	for _, file := range dirFiles {
		if strings.HasPrefix(file.Name(), ".") {
			continue // skip hidden files
		}
		name := dirPath + "/" + file.Name()
		fmt.Println(name)
		if file.IsDir() {
			foundDirs = append(foundDirs, name)
		}
	}
	return foundDirs
}

func breadthFirst(f func(filepath string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		files := worklist
		worklist = nil
		for _, file := range files {
			if !seen[file] {
				seen[file] = true
				worklist = append(worklist, f(file)...)
			}
		}
	}
}
