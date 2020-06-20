package enforce

import (
	"fmt"
	"github.com/lucasew/golisp/data"
	"github.com/lucasew/golisp/data/types/number"
)

var ErrNotANumber = fmt.Errorf("not a number")

func Number(d []data.LispValue, nth int) func() error {
	return func() error {
		v := d[nth-1]
		if !number.IsNumber(v) {
			return fmt.Errorf("%d nth parameter expects a number, got %s", nth, v.LispTypeName())
		}
		return nil
	}
}
