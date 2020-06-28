// Using a select statement, add a timeout to the echo server from Section 8.3 so that it
// disconnects any client that shouts nothing within 10 seconds.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()

	shouts := make(chan string)
	done := make(chan struct{}, 1)
	go func() {
		input := bufio.NewScanner(c)
		for input.Scan() {
			shouts <- input.Text()
		}
		switch err := input.Err(); err.(type) {
		case nil, *net.OpError:
			// ignore net.OpError caused by timeout closing conn before scan loop ends
		default:
			log.Println(err)
		}
		done <- struct{}{} // buffer prevents goroutine leak when handleConn returns first
	}()

	var wg sync.WaitGroup
	for {
		select {
		case <-time.After(10 * time.Second):
			log.Println("client disconnected; connection timed out")
			return
		case shout := <-shouts:
			wg.Add(1)
			go func() {
				echo(c, shout, 1*time.Second)
				wg.Done()
			}()
		case <-done:
			wg.Wait()
			return
		}
	}
}

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}
