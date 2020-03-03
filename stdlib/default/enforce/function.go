package enforce

import (
    "fmt"
    "github.com/lucasew/golisp/data"
    "github.com/lucasew/golisp/data/types"
)

func Function(d []data.LispValue, nth int) func()error {
    return func()error {
        v := d[nth - 1]
        if !types.IsFunction(v) {
            return fmt.Errorf("%d nth parameter expects a function, got %s", nth, v.LispTypeName())
        }
        return nil
    }
}
