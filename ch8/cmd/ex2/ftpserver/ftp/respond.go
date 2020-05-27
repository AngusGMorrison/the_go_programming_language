package ftp

import (
	"fmt"
	"log"
)

// respond copies a string to the client and terminates it with the appropriate FTP line terminator
// for the datatype.
func (c *Conn) respond(s string) {
	log.Print(">> ", s)
	_, err := fmt.Fprint(c.conn, s, c.EOL())
	if err != nil {
		log.Print(err)
	}
}

// EOL returns the line terminator matching the FTP standard for the datatype.
func (c *Conn) EOL() string {
	switch c.dataType {
	case ascii:
		return "\r\n"
	case binary:
		return "\n"
	default:
		return "\n"
	}
}
