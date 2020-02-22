package test

import (
    "testing"
    "github.com/lucasew/golisp/data"
    "github.com/lucasew/golisp/data/is"
)

func String(v data.LispValue) func(*testing.T) {
    return func(t *testing.T) {
        if !is.String(v) {
            t.Fail()
        }
    }
}
