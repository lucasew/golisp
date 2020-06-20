package enforce

import (
	"github.com/lucasew/golisp/data"
)

func Portal(d []data.LispValue, nth int) func() error {
	return Entity("portal", d, nth)
}
