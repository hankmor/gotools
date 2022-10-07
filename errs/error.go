package errs

import "fmt"

func Throw(err error) {
	if err != nil {
		panic(err)
	}
}

func Throwf(err error, format string, v ...any) {
	if err != nil {
		panic(fmt.Sprintf(format, v...))
	}
}
