package types

import (
    "github.com/lucasew/golisp/data"
    "testing"
)

type Symbol string

func NewSymbol(s string) data.LispValue {
    return Symbol(s)
}

func (s Symbol) ToString() string {
    return string(s)
}

func (Symbol) IsNil() bool {
    return false
}

func (s Symbol) Repr() string {
    return string(s)
}

func (Symbol) LispTypeName() string {
    return "symbol"
}

func IsSymbol(v data.LispValue) bool {
    _, ok := v.(Symbol)
    return ok
}

func SymbolTest(v data.LispValue) func(*testing.T) {
    return func(t *testing.T) {
        if !IsSymbol(v) {
            t.Fail()
        }
    }
}
