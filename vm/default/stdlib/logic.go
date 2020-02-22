package stdlib

import (
    "github.com/lucasew/golisp/data"
    "github.com/lucasew/golisp/data/types"
    "github.com/lucasew/golisp/data/convert"
    "github.com/lucasew/golisp/vm/default/stdlib/enforce"
)

func init() {
    register("not", Not)
    register("and", And)
    register("or", Or)
}

func Not(v data.LispCons) (data.LispValue, error) {
    err := enforce.Length(v, 1)
    if err != nil {
        return types.Nil, err
    }
    return convert.NewLispValue(!v.Car().IsNil())
}

func And(v data.LispCons) (data.LispValue, error) {
    if v.Car().IsNil() {
        return types.Nil, nil
    }
    if v.Cdr().IsNil() {
        return v.Car(), nil
    }
    return And(v.Cdr().(data.LispCons))
}

func Or(v data.LispCons) (data.LispValue, error) {
    if v.IsNil() {
        return types.Nil, nil
    }
    if !v.Car().IsNil() {
        return v.Car(), nil
    }
    return Or(v.Cdr().(data.LispCons))
}
