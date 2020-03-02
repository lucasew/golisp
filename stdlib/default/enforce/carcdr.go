package enforce

import (
	"fmt"
	"github.com/lucasew/golisp/data"
	"github.com/lucasew/golisp/data/types/test"
)

func CarCdr(d data.LispValue, nth int) error {
	if !test.IsCarCdr(d) {
		return fmt.Errorf("%d nth parameter expects a carcdr, got %s", nth, d.LispTypeName())
	}
	return nil
}