package enforce

import (
	"fmt"
	"github.com/lucasew/golisp/data"
)

var ErrNotANumber = fmt.Errorf("not a number")

func Number(d []data.LispValue, nth int) func() error {
	return Entity("lisp_number", d, nth)
}
