package enforce

import (
    "github.com/lucasew/golisp/data"
    "github.com/lucasew/golisp/data/types"
    "fmt"
)

var ErrNotANumber = fmt.Errorf("not a number")

func Number(v data.LispValue) error {
    if !types.IsNumber(v) {
        return ErrNotANumber
    }
    return nil
}