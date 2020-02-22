package enforce

import (
    "github.com/lucasew/golisp/data"
    "github.com/lucasew/golisp/data/is"
    "fmt"
)

var ErrNotANumber = fmt.Errorf("not a number")

func Number(v data.LispValue) error {
    if !is.Number(v) {
        return ErrNotANumber
    }
    return nil
}
