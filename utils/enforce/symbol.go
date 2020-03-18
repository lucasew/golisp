package enforce

import (
	"github.com/lucasew/golisp/data"
)

func Symbol(d []data.LispValue, nth int) func() error {
	return Entity("symbol", d, nth)
}
