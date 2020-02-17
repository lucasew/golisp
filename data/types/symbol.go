package types

import (
    "github.com/lucasew/golisp/data"
)

type Symbol string

func NewSymbol(s string) data.LispValue {
    return Symbol(s)
}

func (s Symbol) ToString() string {
    return string(s)
}

func (Symbol) Car() data.LispValue {
    return data.Nil
}

func (Symbol) Cdr() data.LispValue {
    return data.Nil
}

func (Symbol) IsNil() bool {
    return false
}

func (s Symbol) Repr() string {
    return string(s)
}
