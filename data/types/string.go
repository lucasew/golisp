package types

import (
    "fmt"
    "github.com/lucasew/golisp/data"
)

type ConventionalString string

func NewConventionalString(s string) data.LispString {
    return ConventionalString(s)
}

func (c ConventionalString) IsNil() bool {
    return len(c) == 0
}

func (c ConventionalString) Car() data.LispValue {
    if len(c) == 0 {
        return data.Nil
    }
    s := string(c)[0]
    return NewByte(s)
}

func (c ConventionalString) Cdr() data.LispCarCdr {
    if len(c) < 2 {
        return data.Nil
    }
    s := string(c)[1:len(c)] // TODO: Testar se não teremos acessos errados aqui
    return NewConventionalString(s)
}

func (c ConventionalString) Repr() string {
    return fmt.Sprintf("\"%s\"", c)
}

func (c ConventionalString) ToString() string {
    return string(c)
}
