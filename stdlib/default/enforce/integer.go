package enforce

import (
    "fmt"
    "github.com/lucasew/golisp/data"
    "github.com/lucasew/golisp/data/types/number"
)

func Integer(v data.LispValue, nth int) func()error {
    return func()error {
        if !number.IsInt(v) {
            return fmt.Errorf("%d nth parameter expects a int, got %s", nth, v.LispTypeName())
        }
        return nil
    }
}
