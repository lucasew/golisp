package vm

import (
    "github.com/lucasew/golisp/data"
)

type LispVM interface {
    Eval(data.LispValue) (data.LispValue, error)
    EnvGet(key string) data.LispValue // if not exist return nil
    EnvSet(key string, value data.LispValue) data.LispValue
}
