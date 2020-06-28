// Change the chat server's network protocol so that each client provides its name on entering.
// Use the name instead of the network address when prefixing each message with its sender's
// identity.
package main

import (
	"log"
	"net"
	"strings"

	"github.com/angusgmorrison/the_go_programming_language/ch8/ex15/client"
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

var (
	entering = make(chan *client.Client)
	leaving  = make(chan *client.Client)
	messages = make(chan string) // all incoming client messages
)

func broadcaster() {
	clients := make(map[*client.Client]bool) // all connected clients
	for {
		select {
		case msg := <-messages:
			// Broadcast incoming message to all clients' outgoing message channels
			for c := range clients {
				select {
				case c.Ch <- msg:
					// Successfully sent; do nothing
				default:
					// Skip message if buffer is full
				}
			}
		case c := <-entering:
			msg := fmtConnectedClients(clients)
			c.Ch <- msg // Send connected list to client before it is included in the list
			clients[c] = true
		case c := <-leaving:
			delete(clients, c)
		}
	}
}

func fmtConnectedClients(clients map[*client.Client]bool) string {
	var sb strings.Builder
	sb.WriteString("Chatting with:\n")
	for c := range clients {
		sb.WriteString(c.Username())
		sb.WriteByte('\n')
	}
	return sb.String()
}

func handleConn(conn net.Conn) {
	c := client.NewClient(conn)
	messages <- c.Username() + " has arrived"
	entering <- c

	for {
		msg, ok := c.ReadInput()
		if !ok {
			break
		}
		messages <- c.Username() + ": " + msg
	}

	leaving <- c
	messages <- c.Username() + " has left"
	c.Close()
}
