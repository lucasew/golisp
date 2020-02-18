package stdlib

import (
    "github.com/lucasew/golisp/data"
    "github.com/lucasew/golisp/data/types"
    "github.com/lucasew/golisp/vm"
    "errors"
)

func init() {
    register("setg", Setg)
}

func Setg(env vm.LispVM, v data.LispValue) (data.LispValue, error) {
    key, ok := v.Car().(types.Symbol)
    if !ok {
        return data.Nil, errors.New("invalid output for setg")
    }
    value, err := env.Eval(v.Cdr().Car())
    if err != nil {
        return data.Nil, err
    }
    env.EnvSetGlobal(key.ToString(), value)
    return value, nil
}
