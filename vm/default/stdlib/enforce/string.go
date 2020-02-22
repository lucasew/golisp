package enforce

import (
    "github.com/lucasew/golisp/data"
    "github.com/lucasew/golisp/data/types"
    "fmt"
)

func String(v data.LispValue, nth int) error {
    if !types.IsString(v) {
        return fmt.Errorf("%d nth element is not a string", nth)
    }
    return nil
}
