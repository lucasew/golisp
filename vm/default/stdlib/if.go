package stdlib

import (
    "github.com/lucasew/golisp/data"
    "github.com/lucasew/golisp/vm"
    "github.com/lucasew/golisp/vm/default/stdlib/enforce"
)

func init() {
    register("if", If)
}

func If(env vm.LispVM, v data.LispCons) (data.LispValue, error) {
    err := enforce.Length(v, 3)
    if err != nil {
        return data.Nil, err
    }
    cond, err := env.Eval(v.Car())
    if err != nil {
        return data.Nil, err
    }
    if !cond.IsNil() {
        return env.Eval(v.Cdr().Car())
    } else {
        return env.Eval(v.Cdr().Cdr().Car())
    }
}
