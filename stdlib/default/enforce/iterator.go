package enforce

import (
	"fmt"
	"github.com/lucasew/golisp/data"
	"github.com/lucasew/golisp/data/types/test"
)

func Iterator(d data.LispValue, nth int) error {
	if !test.IsIterator(d) {
		return fmt.Errorf("%d nth parameter expects a iterator, got %s", nth, d.LispTypeName())
	}
	return nil
}
