package enforce

import (
    "github.com/lucasew/golisp/data"
    "fmt"
)

func Cons(d data.LispValue, nth int) error {
    _, ok := d.(data.LispCons)
    if !ok {
        return fmt.Errorf("%d nth parameter expects a cons, got %T", nth, d)
    }
    return nil
}
