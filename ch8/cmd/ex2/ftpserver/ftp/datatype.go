package ftp

type dataType int

const (
	ascii dataType = iota
	binary
)

func (c *Conn) setDataType(args []string) {
	if len(args) == 0 {
		c.respond(status501)
	}

	switch args[0] {
	case "A":
		c.dataType = ascii
	case "I": // image/binary
		c.dataType = binary
	default:
		c.respond(status504)
		return
	}
	c.respond(status200)
}
