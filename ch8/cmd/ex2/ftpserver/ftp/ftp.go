// Package ftp provides structs and functions for running an FTP server.
package ftp

import (
	"bufio"
	"log"
	"strings"
)

const (
	status150 = "150 File status okay; about to open data connection."
	status200 = "200 Command okay."
	status220 = "220 Service ready for new user."
	status221 = "221 Service closing control connection."
	status226 = "226 Closing data connection. Requested file action successful."
	status230 = "230 User %s logged in, proceed."
	status425 = "425 Can't open data connection."
	status426 = "426 Connection closed; transfer aborted."
	status501 = "501 Syntax error in parameters or arguments."
	status502 = "502 Command not implemented."
	status504 = "504 Cammand not implemented for that parameter."
	status550 = "550 Requested action not taken. File unavailable."
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
}
