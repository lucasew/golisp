package enforce

import (
	"fmt"
	"github.com/lucasew/golisp/data"
	"github.com/lucasew/golisp/data/types/test"
)

func Iterator(d []data.LispValue, nth int) func() error {
	return func() error {
		v := d[nth-1]
		if !test.IsIterator(v) {
			return fmt.Errorf("%d nth parameter expects a iterator, got %s", nth, v.LispTypeName())
		}
		return nil
	}
}
