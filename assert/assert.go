package assert

import (
	"fmt"
)

func Require(cond bool, format string, v ...any) {
	if !cond {
		panic(fmt.Sprintf(format, v...))
	}
}
