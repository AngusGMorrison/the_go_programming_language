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
	args := make([]string, len(c.args))
	for i, arg := range c.args {
		args[i] = arg.String()
	}
	return fmt.Sprintf("%s(%s)", c.fn, strings.Join(args, ", "))
}

func (m min) String() string {
	items := make([]string, len(m))
	for i, item := range m {
		items[i] = item.String()
	}
	return fmt.Sprintf("min(%s)", strings.Join(items, ", "))
}
