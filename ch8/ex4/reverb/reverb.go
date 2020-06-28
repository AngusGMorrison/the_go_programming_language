// Modify the reverb2 server to use a sync.WaitGroup per connection to count the number of active
// echo goroutines. When it falls to zero, close the write half of the TCP connection as described
// in Exercise 8.3. Verify that your modified netcat3 client from that exercise waits for the final
// echoes of multiple concurrent shouts, even after the standard input has been closed.
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
			log.Print(err)
			continue
		}
		handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()

	input := bufio.NewScanner(c)
	var wg sync.WaitGroup
	for input.Scan() {
		if input.Err() != nil {
			log.Print(input.Err())
			return
		}
		wg.Add(1)
		go func(shout string) {
			defer wg.Done()
			echo(c, shout, 1*time.Second)
		}(input.Text())
	}

	wg.Wait()
}

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}
