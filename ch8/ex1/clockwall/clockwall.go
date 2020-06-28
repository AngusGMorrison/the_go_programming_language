// Modify clock2 to accept a port number, and write a program, clockwall, that acts as a client of
// several clock servers at once, reading the times from each one and displaying the results in a
// table, akin to the wall of clocks seen in some business offices.
package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"sort"
	"strings"
	"time"
)

func main() {
	servers, err := parse(os.Args[1:])
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", os.Args[0], err)
		os.Exit(1)
	}
	sort.Sort(servers) // alphabetize clock wall

	conns, err := connect(servers)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", os.Args[0], err)
	}
	for _, conn := range conns {
		defer conn.Close()
	}

	displayClocks(servers)
}

func parse(args []string) (cw clockWall, err error) {
	if len(args) == 0 {
		return nil, fmt.Errorf("usage: %s Location=address:port [...]", os.Args[0])
	}
	for _, arg := range args {
		sDetails := strings.SplitN(arg, "=", 2)
		if len(sDetails) != 2 {
			return nil, fmt.Errorf("%s doesn't match format 'name=address'", arg)
		}
		cw = append(cw, &server{sDetails[0], sDetails[1], ""})
	}
	return
}

func connect(cw clockWall) ([]net.Conn, error) {
	conns := make([]net.Conn, 0, len(cw))
	for _, server := range cw {
		conn, err := net.Dial("tcp", server.address) // connect to the server
		if err != nil {
			return conns, fmt.Errorf("couldn't connect to %s", server.address)
		}
		conns = append(conns, conn)
		go mustCopy(server, conn)
	}
	return conns, nil
}

func mustCopy(server *server, conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		server.output = scanner.Text()
		if err := scanner.Err(); err != nil {
			fmt.Fprintf(os.Stderr, "failed to scan %s: %v", server.address, err)
			os.Exit(1)
		}
	}
}

func displayClocks(cw clockWall) {
	fmt.Printf("%s\n", cw.String())
	for {
		for _, server := range cw {
			fmt.Print(server.String())
		}
		fmt.Println()
		time.Sleep(1 * time.Second)
	}
}
