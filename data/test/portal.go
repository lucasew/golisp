package test

import (
    "testing"
    "github.com/lucasew/golisp/data"
    "github.com/lucasew/golisp/data/is"
)

func Portal(v data.LispValue) func(t *testing.T) {
    return func(t *testing.T) {
        if !is.Portal(v) {
            t.Fail()
        }
    }
}
