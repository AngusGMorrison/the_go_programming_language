// Implement a concurrent File Transfer Protocol (FTP) server. The server should interpret commands
// from each client such as cd to change directory, ls to list a directory, get to send the contents
// of a file, and close to close the connection.
//
// Thanks to Kdama for the solution that got me going https://github.com/kdama/gopl/blob/master/ch08/ex02/main.go
package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"path/filepath"

	"the_go_programming_language/ch8/ex2/ftpserver/ftp"
)

var port int
var rootDir string

func init() {
	flag.IntVar(&port, "port", 8080, "port number")
	flag.StringVar(&rootDir, "rootDir", "public", "root directory")
	flag.Parse()
}

func main() {
	server := fmt.Sprintf(":%d", port)
	listener, err := net.Listen("tcp", server)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	absPath, err := filepath.Abs(rootDir)
	if err != nil {
		log.Fatal(err)
	}
	ftp.Serve(ftp.NewConn(c, absPath))
}
