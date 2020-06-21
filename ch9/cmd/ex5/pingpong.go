// Write a program with two goroutines that send messages back and forth over two unbuffered
// channels in ping-pong fashion. How many communications per second can the program sustain?
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch1 := make(chan struct{})
	ch2 := make(chan struct{})

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		var msgCount uint64
		ticker := time.Tick(1 * time.Second)
		for {
			select {
			case rec := <-ch1:
				msgCount++
				ch2 <- rec
			case <-ticker:
				fmt.Printf("%d msg/s\n", msgCount*2)
				msgCount = 0
			}
		}
	}()

	go func() {
		for {
			ch1 <- <-ch2
		}
	}()

	ch1 <- struct{}{} // Start an infinite message cycle
	wg.Wait()
}
