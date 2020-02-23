package enforce

import (
	"fmt"
	"github.com/lucasew/golisp/data"
	"github.com/lucasew/golisp/data/types/iterator"
)

func Iterator(d data.LispValue, nth int) error {
	if !iterator.IsIterator(d) {
		return fmt.Errorf("%d nth parameter expects a iterator, got %s", nth, d.LispTypeName())
	}
	return nil
}
