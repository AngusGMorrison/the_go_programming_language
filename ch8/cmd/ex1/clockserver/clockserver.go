// clockserver is a TCP server that periodically writes the time.
package main

import (
	"flag"
	"io"
	"log"
	"net"
	"time"
)

var port = flag.String("port", "8000", "port")

func main() {
	flag.Parse()
	address := "localhost:" + *port
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return // e.g., client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}
