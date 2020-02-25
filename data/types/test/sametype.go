package test

import (
    "reflect"
    "github.com/lucasew/golisp/data"
)

func IsSameType(a, b data.LispValue) bool {
    return reflect.TypeOf(a) == reflect.TypeOf(b)
}
