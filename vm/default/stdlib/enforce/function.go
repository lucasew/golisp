package enforce

import (
    "github.com/lucasew/golisp/data"
    "fmt"
)

func Function(d data.LispValue, nth int) error {
    _, ok := d.(data.LispFunction)
    if !ok {
        return fmt.Errorf("%d nth parameter expects a function, got %T", nth, d)
    }
    return nil
}
