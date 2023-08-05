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
	return fmt.Sprintf("%s%s", string(u.op), u.x)
}

func (b binary) String() string {
	return fmt.Sprintf("%s %s %s", b.x, string(b.op), b.y)
}

func (c call) String() string {
	var strArgs []string
	for _, a := range c.args {
		strArgs = append(strArgs, a.String())
	}

	return fmt.Sprintf("%s(%s)", c.fn, strings.Join(strArgs, ", "))
}
