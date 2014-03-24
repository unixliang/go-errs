package errs

import (
	"fmt"
	"path"
	"runtime"
)

type Error struct {
	msg string
	st  string
}

// Error with stace
func New(msg string) error {
	return &Error{msg, stack(3)}

}

func Newf(format string, args ...interface{}) error {
	return &Error{fmt.Sprintf(format, args ...), stack(3)}
}

func (e *Error) Error() string {
	return e.msg + e.st
}

func stack(skip int) string {
	stk := make([]uintptr, 32)
	str := ""
	l := runtime.Callers(skip, stk[:])
	for i := 0; i < l; i++ {
		f := runtime.FuncForPC(stk[i])
		name := f.Name()
		file, line := f.FileLine(stk[i])
		str += fmt.Sprintf("\n    %-30s [%s:%d]", name, path.Base(file), line)
	}
	return str
}
