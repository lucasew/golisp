package vm

import (
    "github.com/lucasew/golisp/datatypes"
)

type LispEnvKV struct {
    Key string
    Value datatypes.LispValue
}

func NewKV(k string, v datatypes.LispValue) LispEnvKV {
    return LispEnvKV{k, v}
}
