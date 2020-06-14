// Package client provides the struct and methods required to model concurrent client connections
// to a chat server.
package client

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

// Client represents a single client connection to the server.
type Client struct {
	username string
	conn     net.Conn
	Ch       chan string
	idle     *time.Timer
	input    *bufio.Scanner
}

func (c *Client) Username() string {
	return c.username
}

var timeout = 5 * time.Minute // how long idle clients remain connected

// SetTimeout sets the duration after which idle clients will be disconnected.
func SetTimeout(t time.Duration) {
	timeout = t
}

// NewClient returns a fully configured client that listens for messages on its channel and writes
// them to its TCP connection.
func NewClient(conn net.Conn) *Client {
	c := &Client{
		conn:  conn,
		Ch:    make(chan string),
		input: bufio.NewScanner(conn),
	}
	c.idle = time.AfterFunc(timeout, func() { c.Close() })
	go c.writer()
	c.setUsername()
	return c
}

// writer listens for outgoing messages on the client's channel and writes them to the connection.
func (c *Client) writer() {
	for msg := range c.Ch {
		fmt.Fprintln(c.conn, msg) // NOTE: ignoring network errors
	}
}

// setUsername prompts the user to enter a name to identify themselves in chat and saves it to the
// client.
func (c *Client) setUsername() {
	c.Ch <- "Enter your name: "
	c.input.Scan()
	c.idle.Stop() // reset idle timer on name input
	c.idle.Reset(timeout)
	c.username = c.input.Text()
}

// ReadInput scans the client connection for input and returns it, resetting the idle timeout to
// maintain the client connection.
func (c *Client) ReadInput() (string, bool) {
	ok := c.input.Scan()
	if !ok {
		return "", false
	}
	c.idle.Stop()
	c.idle.Reset(timeout)
	return c.input.Text(), true
}

// Close stops all of client's channels and associated goroutines before closing the connection.
func (c *Client) Close() {
	c.idle.Stop()
	close(c.Ch)
	c.conn.Close()
}
