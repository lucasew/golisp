package enforce

import (
    "github.com/lucasew/golisp/data"
    "github.com/lucasew/golisp/data/is"
    "fmt"
)

func Function(d data.LispValue, nth int) error {
    if !is.Function(d) {
        return fmt.Errorf("%d nth parameter expects a function, got %T", nth, d)
    }
    return nil
}
