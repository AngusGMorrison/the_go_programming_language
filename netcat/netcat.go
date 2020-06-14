// In netcat, the interface value conn has the concrete type *net.TCPConn, whihc represents a TCP
// connection. A TCP connection constists of two halves that may be closed independently using its
// CloseRead and CloseWrite methods. Modify the main goroutine of netcat3 to close only the write
// half of the connection so that the program will continue to print the final echoes from the
// reverb1 server even after the standard input has been closed. (Doing this for the reverb2 server
// is harder; see Exercise 8.4.)
package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn)
		log.Println("done")
		done <- struct{}{} // signal the main goroutine
	}()
	mustCopy(conn, os.Stdin)
	conn.(*net.TCPConn).CloseWrite()
	<-done // wait for background goroutine to finish
}

func mustCopy(dest io.Writer, src io.Reader) {
	_, err := io.Copy(dest, src)
	if err != nil {
		log.Fatal(err)
	}
}
