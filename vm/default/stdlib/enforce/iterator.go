
package enforce

import (
    "github.com/lucasew/golisp/data"
    "github.com/lucasew/golisp/data/types/iterator"
    "fmt"
)

func Iterator(d data.LispValue, nth int) error {
    if !iterator.IsIterator(d) {
        return fmt.Errorf("%d nth parameter expects a iterator, got %s", nth, d.LispTypeName())
    }
    return nil
}
