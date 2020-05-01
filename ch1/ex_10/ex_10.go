/*
Find a website that produces a large amount of data. Investigate caching by
running fetchall twice in succession to see whether the reported time changes
much. Do you get the same content each time? Modify fetchall to print its output
to a file so it can be examined.
*/

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}

	outputFilename := "out.json"
	output, err := os.Create(outputFilename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "main: couldn't open %s\n", outputFilename)
		os.Exit(1)
	}

	for range os.Args[1:] {
		fmt.Fprintf(output, <-ch) // Receive from channel ch
	}
	output.Close()
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %s\n%s\n\n", secs, url, body)
}
