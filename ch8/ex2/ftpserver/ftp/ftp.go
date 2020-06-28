// Package ftp provides structs and functions for running an FTP server.
package ftp

import (
	"bufio"
	"log"
	"strings"
)

// Serve scans incoming requests for valid commands and routes them to handler functions.
func Serve(c *Conn) {
	c.respond(status220)

	s := bufio.NewScanner(c.conn)
	for s.Scan() {
		input := strings.Fields(s.Text())
		if len(input) == 0 {
			continue
		}
		command, args := input[0], input[1:]
		log.Printf("<< %s %v", command, args)

		switch command {
		case "CWD": // cd
			c.cwd(args)
		case "LIST": // ls
			c.list(args)
		case "PORT":
			c.port(args)
		case "USER":
			c.user(args)
		case "QUIT": // close
			c.respond(status221)
			return
		case "RETR": // get
			c.retr(args)
		case "TYPE":
			c.setDataType(args)
		default:
			c.respond(status502)
		}
	}
	if s.Err() != nil {
		log.Print(s.Err())
	}
}
