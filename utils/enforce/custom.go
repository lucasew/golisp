package enforce

import (
	"fmt"
	"github.com/lucasew/golisp/data"
)

func Custom(d []data.LispValue, nth int, isFunc func(v data.LispValue) bool, typename string) func() error {
	return func() error {
		v := d[nth-1]
		if !isFunc(v) {
			return fmt.Errorf("%d nth parameter expects a %s, got %s", nth, typename, v.LispTypeName())
		}
		return nil
	}
}
