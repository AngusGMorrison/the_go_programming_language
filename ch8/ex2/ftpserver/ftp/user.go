package ftp

import (
	"fmt"
	"strings"
)

func (c *Conn) user(args []string) {
	c.respond(fmt.Sprintf(status230, strings.Join(args, " ")))
}
