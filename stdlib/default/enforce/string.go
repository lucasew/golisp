package enforce

import (
    "fmt"
    "github.com/lucasew/golisp/data"
    "github.com/lucasew/golisp/data/types"
)

func String(v data.LispValue, nth int) func()error {
    return func()error {
        if !types.IsString(v) {
            return fmt.Errorf("%d nth element is not a string", nth)
        }
        return nil
    }
}
