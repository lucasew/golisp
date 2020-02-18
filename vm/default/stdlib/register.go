package stdlib

import (
    "github.com/lucasew/golisp/data"
    "github.com/lucasew/golisp/data/convert"
    "github.com/lucasew/golisp/vm/components/env"
)

var stdlib = map[string]data.LispValue{}

func register(k string, v interface{}) {
    var err error
    stdlib[k], err = convert.NewLispValue(v)
    if err != nil {
        panic(err)
    }
}

func NewDefaultEnv(parent *env.LispEnv) *env.LispEnv {
    parent = env.NewLispEnv(parent)
    for k, v := range stdlib {
        parent.SetLocal(k, v)
    }
    return parent
}
