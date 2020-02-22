package is

import (
    "github.com/lucasew/golisp/data"
    "reflect"
)

func SameType(a, b data.LispValue) bool {
    return reflect.TypeOf(a) == reflect.TypeOf(b)
}
