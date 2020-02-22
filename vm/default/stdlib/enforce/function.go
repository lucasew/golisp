package enforce

import (
    "github.com/lucasew/golisp/data"
    "github.com/lucasew/golisp/data/types"
    "fmt"
)

func Function(d data.LispValue, nth int) error {
    if !types.IsFunction(d) {
        return fmt.Errorf("%d nth parameter expects a function, got %s", nth, d.LispTypeName())
    }
    return nil
}
