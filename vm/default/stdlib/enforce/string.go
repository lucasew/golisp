package enforce

import (
    "github.com/lucasew/golisp/data"
    "fmt"
)

func String(v data.LispValue, nth int) error {
    _, ok := v.(data.LispString)
    if !ok {
        return fmt.Errorf("%d nth element is not a string", nth)
    }
    return nil
}
