package enforce

import (
    "github.com/lucasew/golisp/data"
    "github.com/lucasew/golisp/data/types"
    "fmt"
)

func Cons(d data.LispValue, nth int) error {
    if !types.IsCons(d) {
        return fmt.Errorf("%d nth parameter expects a cons, got %T", nth, d)
    }
    return nil
}
