package test

import (
    "testing"
    "github.com/lucasew/golisp/data"
    "github.com/lucasew/golisp/data/is"
)

func Map(v data.LispValue) func(t *testing.T) {
    return func(t *testing.T) {
        if !is.Map(v) {
            t.Fail()
        }
    }
}
