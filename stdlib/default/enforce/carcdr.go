package enforce

import (
    "fmt"
    "github.com/lucasew/golisp/data"
    "github.com/lucasew/golisp/data/types/test"
)

func CarCdr(d []data.LispValue, nth int) func()error {
    return func()error {
        v := d[nth - 1]
        if !test.IsCarCdr(v) {
            return fmt.Errorf("%d nth parameter expects a carcdr, got %s", nth, v.LispTypeName())
        }
        return nil
    }
}
