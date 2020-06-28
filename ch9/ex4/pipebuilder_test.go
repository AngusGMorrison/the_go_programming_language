package pipebuilder

import (
	"fmt"
	"testing"
	"time"
)

func TestBuild(t *testing.T) {
	start := make(chan struct{})
	end := make(chan struct{})
	stages := 1000
	go build(start, end, stages)

	<-end // wait for ready signal
	begin := time.Now()
	start <- struct{}{} // time transit from start to end
	<-end
	fmt.Printf("Took %s with %d stages\n", time.Since(begin), stages)
}
