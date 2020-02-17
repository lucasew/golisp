package env

import (
    "github.com/lucasew/golisp/data"
)

type LispEnvKV struct {
    Key string
    Value data.LispValue
}

func NewKV(k string, v data.LispValue) LispEnvKV {
    return LispEnvKV{k, v}
}
