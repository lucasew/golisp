package stdlib

import (
    "github.com/lucasew/golisp/data"
    "github.com/lucasew/golisp/data/convert"
)

func init() {
    register("not", Not)
    register("and", And)
    register("or", Or)
}

func Not(v data.LispValue) (data.LispValue, error) {
    return convert.NewLispValue(v.Car().IsNil())
}

func And(v data.LispValue) (data.LispValue, error) {
    if v.Car().IsNil() {
        return data.Nil, nil
    }
    if v.Cdr().IsNil() {
        return v.Car(), nil
    }
    return And(v.Cdr())
}

func Or(v data.LispValue) (data.LispValue, error) {
    if v.IsNil() {
        return data.Nil, nil
    }
    if !v.Car().IsNil() {
        return v.Car(), nil
    }
    return Or(v.Cdr())
}
