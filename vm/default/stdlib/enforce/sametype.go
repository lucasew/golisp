package enforce

import (
    "github.com/lucasew/golisp/data"
    "reflect"
    "fmt"
)

var ErrNotSameType = fmt.Errorf("not the same type")

func SameType(a, b data.LispValue) error {
    ta := reflect.TypeOf(a)
    tb := reflect.TypeOf(b)
    if ta != tb {
        return fmt.Errorf("%w: expected %T got %T", ErrNotSameType, a, b)
    }
    return nil
}