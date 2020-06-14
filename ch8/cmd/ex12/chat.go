// Make the broadcaster announce the current set of clients to each new arrival. This requires
// that the clients set and entering and leaving channels record the client name too.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

// A named outgoing message channel.
type client struct {
	name string
	ch   chan string
}

var (
	entering = make(chan *client)
	leaving  = make(chan *client)
	messages = make(chan string) // all incoming client messages
)

func broadcaster() {
	clients := make(map[*client]bool) // all connected clients
	for {
		select {
		case msg := <-messages:
			// Broadcast incoming message to all clients' outgoing message channels
			for cli := range clients {
				cli.ch <- msg
			}
		case cli := <-entering:
			msg := fmtConnectedClients(clients)
			cli.ch <- msg // Send connected list to client before it is included in the list
			clients[cli] = true
		case cli := <-leaving:
			delete(clients, cli)
			close(cli.ch)
		}
	}
}

func fmtConnectedClients(clients map[*client]bool) string {
	var sb strings.Builder
	sb.WriteString("Chatting with:\n")
	for cli := range clients {
		sb.WriteString(cli.name)
		sb.WriteByte('\n')
	}
	sb.WriteByte('\n')
	return sb.String()
}

func handleConn(conn net.Conn) {
	cli := &client{
		name: conn.RemoteAddr().String(),
		ch:   make(chan string),
	}
	go clientWriter(conn, cli.ch)
	cli.ch <- "You are " + cli.name
	messages <- cli.name + " has arrived"
	entering <- cli

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- cli.name + ": " + input.Text()
	}
	// NOTE: ignoring potential errors from input.Err()

	leaving <- cli
	messages <- cli.name + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // NOTE: ignoring network errors
	}
}
