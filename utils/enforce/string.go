package enforce

import (
	"github.com/lucasew/golisp/data"
)

func String(d []data.LispValue, nth int) func() error {
	return Entity("lisp_string", d, nth)
}
