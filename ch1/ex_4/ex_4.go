package main

// Modify dup2 to print the names aof all files in which each duplicated line
// occurs.

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	inFiles := make(map[string][]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, inFiles)
	} else {
		for _, arg := range files {
			file, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(file, counts, inFiles)
			file.Close()
		}
	}
	for line, count := range counts {
		if count > 1 {
			fmt.Printf("%v:\n", strings.Join(inFiles[line], ", "))
			fmt.Printf("%d\t%s\n", count, line)
		}
	}
}

func countLines(file *os.File, counts map[string]int, inFiles map[string][]string) {
	input := bufio.NewScanner(file)
	filename := file.Name()
	for input.Scan() {
		text := input.Text()
		counts[text]++
		if !fileAlreadyFound(filename, inFiles[text]) {
			inFiles[text] = append(inFiles[text], filename)
		}
	}
	// NOTE: ignoring potential errors from input.Err()
}

func fileAlreadyFound(filename string, foundFiles []string) bool {
	for _, foundName := range foundFiles {
		if filename == foundName {
			return true
		}
	}
	return false
}
