package stdlib

import (
    "github.com/lucasew/golisp/data"
    "github.com/lucasew/golisp/vm"
    "fmt"
)

func init() {
    register("quote", Quote)
    register("list", List)
    register("call", Call)
}

func Quote(env vm.LispVM, v data.LispCons) (data.LispValue, error) {
    return v.Car(), nil
}

func List(v data.LispCons) (data.LispValue, error) {
    return v, nil
}

func Call(v data.LispCons) (data.LispValue, error) {
    fn, ok := v.Car().(data.LispFunction)
    if !ok {
        return data.Nil, fmt.Errorf("invalid first value. expected function got %T", v.Car())
    }
    val, ok := v.Cdr().Car().(data.LispCons)
    if !ok {
        return data.Nil, fmt.Errorf("invalid second value. expected cons got %T", v.Car())
    }
    return fn.LispCall(val)
}
