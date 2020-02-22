package types

import (
    "github.com/lucasew/golisp/data"
    "testing"
)

func IsNumber(v data.LispValue) bool {
    _, ok := v.(data.LispNumber)
    return ok
}


func NumberTest(v data.LispValue) func(t *testing.T) {
    return func(t *testing.T) {
        if !IsNumber(v) {
            t.Fail()
        }
    }
}
