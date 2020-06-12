// Write a version of du that computes and periodically displays separate totals for each of the
// root directories.
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

// file holds the size of a file in bytes and the root dir of the subtree it belongs to.
type file struct {
	root string
	size int64
}

// rootInfo holds the name, file count and total numberof bytes of a subtree root directory.
type rootInfo struct {
	name           string
	nfiles, nbytes int64
}

var verbose = flag.Bool("v", false, "show verbose progress messages")

func main() {
	// Determine the initial directories and begin concurrent traversal.
	flag.Parse()
	args := flag.Args()
	roots := make(map[string]*rootInfo) // track each root separately in a map
	if len(args) == 0 {
		args = append(args, ".")
	}
	files := make(chan *file)
	var wg sync.WaitGroup
	for _, arg := range args {
		roots[arg] = &rootInfo{name: arg}
		wg.Add(1)
		go walkDir(arg, arg, &wg, files) // along with current dir, provide root dir to goroutine
	}
	go func() {
		wg.Wait()
		close(files)
	}()

	// Print the results periodically.
	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}
loop:
	for {
		select {
		case file, ok := <-files:
			if !ok {
				break loop // fileSizes was closed
			}
			root := roots[file.root] // assign received size to correct root
			root.nfiles++
			root.nbytes += file.size
		case <-tick:
			printDiskUsage(roots)
		}
	}
	printDiskUsage(roots) // final totals
}

func walkDir(dir, root string, n *sync.WaitGroup, fileSizes chan<- *file) {
	defer n.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(subdir, root, n, fileSizes)
		} else {
			fileSizes <- &file{root, entry.Size()} // file informs receiver of correct root
		}
	}
}

// sema is a counting semaphore for limiting concurrency in dirents.
var sema = make(chan struct{}, 20)

// dirents returns the entries of directory dir.
func dirents(dir string) []os.FileInfo {
	sema <- struct{}{}        // acquire token
	defer func() { <-sema }() // release token

	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", os.Args[0], err)
		return nil
	}
	return entries
}

func printDiskUsage(roots map[string]*rootInfo) {
	var totalFiles, totalBytes int64
	for name, info := range roots {
		totalFiles += info.nfiles
		totalBytes += info.nbytes
		fmt.Printf("%s: %d files  %.1f GB\n", name, info.nfiles, float64(info.nbytes)/1e9)
	}
	fmt.Printf("Total: %d files  %.1f GB\n\n", totalFiles, float64(totalBytes)/1e9)
}
