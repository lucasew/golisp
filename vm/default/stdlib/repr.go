package stdlib

import (
    "github.com/lucasew/golisp/data"
    "github.com/lucasew/golisp/data/types"
)

func init() {
    register("repr", Repr)
}

func Repr(v data.LispCons) (data.LispValue, error) {
    return types.NewConventionalString(v.Car().Repr()), nil
}

