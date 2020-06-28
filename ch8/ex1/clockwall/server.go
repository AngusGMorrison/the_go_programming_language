package main

import (
	"fmt"
	"strings"
)

type server struct {
	name, address, output string
}

func (s server) String() string {
	return fmt.Sprintf("%-12s\t", s.output)
}

// clockWall provides a slice of *server that satisfies sort.Interface
type clockWall []*server

func (c clockWall) Len() int {
	return len(c)
}

func (c clockWall) Less(i, j int) bool {
	return c[i].name < c[j].name
}

func (c clockWall) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func (c clockWall) String() string {
	var sb strings.Builder
	for _, server := range c {
		sb.WriteString(fmt.Sprintf("%-12s\t", server.name)) // construct the table header
	}
	return sb.String()
}
