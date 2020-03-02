package enforce

import (
	"fmt"
	"github.com/lucasew/golisp/data"
	"github.com/lucasew/golisp/data/types/number"
)

var ErrNotANumber = fmt.Errorf("not a number")

func Number(v data.LispValue, nth int) error {
	if !number.IsNumber(v) {
		return fmt.Errorf("%d nth parameter expects a number, got %s", nth, v.LispTypeName())
	}
	return nil
}
