package types

import (
    "reflect"
    "github.com/lucasew/golisp/data"
)

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

var CarCdrTest = NewTestHelper(IsCarCdr)
var LenTest = NewTestHelper(IsLen)
var MapTest = NewTestHelper(IsMap)
var ValueTest = NewTestHelper(IsValue)
