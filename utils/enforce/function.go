package enforce

import (
	"github.com/lucasew/golisp/data"
)

func Function(d []data.LispValue, nth int) func() error {
	return Entity("lisp_function", d, nth)
}
