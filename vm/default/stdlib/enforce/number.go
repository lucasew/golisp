package enforce

import (
    "github.com/lucasew/golisp/data"
    "fmt"
)

var ErrNotANumber = fmt.Errorf("not a number")

func Number(v data.LispValue) error {
    _, ok := v.(data.LispNumber)
    if !ok {
        return ErrNotANumber
    }
    return nil
}
