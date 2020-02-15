package vm

import (
    "github.com/lucasew/golisp/datatypes"
)

type LispVM interface {
    Eval(datatypes.LispValue) (datatypes.LispValue, error)
    EnvGet(key string) datatypes.LispValue // if not exist return nil
    EnvSet(key string, value datatypes.LispValue) datatypes.LispValue
}
