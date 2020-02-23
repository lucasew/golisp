package enforce

import (
	"fmt"
	"github.com/lucasew/golisp/data"
	"github.com/lucasew/golisp/data/types"
)

func Cons(d data.LispValue, nth int) error {
	if !types.IsCons(d) {
		return fmt.Errorf("%d nth parameter expects a cons, got %s", nth, d.LispTypeName())
	}
	return nil
}
