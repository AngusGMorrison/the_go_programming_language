package eval

import (
	"fmt"
	"strings"
)

func (v Var) String() string {
	return string(v)
}

func (l literal) String() string {
	return fmt.Sprintf("%.6g", l)
}

func (u unary) String() string {
	return fmt.Sprintf("%c%s", u.op, u.x.String())
}

func (b binary) String() string {
	return fmt.Sprintf("(%s %c %s)", b.x.String(), b.op, b.y.String())
}

func (c call) String() string {
	args := []string{}
	for _, arg := range c.args {
		args = append(args, arg.String())
	}
	return fmt.Sprintf("%s(%s)", c.fn, strings.Join(args, ", "))
}
