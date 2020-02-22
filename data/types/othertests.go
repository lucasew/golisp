package types

import (
    "testing"
    "reflect"
    "github.com/lucasew/golisp/data"
)

func CarCdrTest(v data.LispValue) func(t *testing.T) {
    return func(t *testing.T) {
        if !IsCarCdr(v) {
            t.Fail()
        }
    }
}

func LenTest(v data.LispValue) func(t *testing.T) {
    return func(t *testing.T) {
        if !IsLen(v) {
            t.Fail()
        }
    }
}

func MapTest(v data.LispValue) func(t *testing.T) {
    return func(t *testing.T) {
        if !IsMap(v) {
            t.Fail()
        }
    }
}

func ValueTest(v data.LispValue) func(t *testing.T) {
    return func(t *testing.T) {
    }
}

func IsCarCdr(v data.LispValue) bool {
    _, ok := v.(data.LispCarCdr)
    return ok
}

func IsLen(v data.LispValue) bool {
    _, ok := v.(data.LispLen)
    return ok
}

func IsMap(v data.LispValue) bool {
    _, ok := v.(data.LispMap)
    return ok
}

func IsSameType(a, b data.LispValue) bool {
    return reflect.TypeOf(a) == reflect.TypeOf(b)
}

func IsValue(v data.LispValue) bool {
    _, ok := v.(data.LispValue)
    return ok
}
