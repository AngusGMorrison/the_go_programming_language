package ftp

import (
	"io"
	"log"
	"os"
	"path/filepath"
)

func (c *Conn) retr(args []string) {
	if len(args) != 1 {
		c.respond(status501)
		return
	}

	path := filepath.Join(c.rootDir, c.workDir, args[0])
	file, err := os.Open(path)
	if err != nil {
		log.Print(err)
		c.respond(status550)
	}
	c.respond(status150)

	dataConn, err := c.dataConnect()
	if err != nil {
		log.Print(err)
		c.respond(status425)
	}
	defer dataConn.Close()

	_, err = io.Copy(dataConn, file)
	if err != nil {
		log.Print(err)
		c.respond(status426)
		return
	}
	io.WriteString(dataConn, c.EOL())
	c.respond(status226)
}
