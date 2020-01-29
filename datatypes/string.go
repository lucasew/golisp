package datatypes

import (
    "fmt"
)

type ConventionalString string

func NewConventionalString(s string) LispString {
    return ConventionalString(s)
}

func (c ConventionalString) IsNil() bool {
    return len(c) == 0
}

func (c ConventionalString) Car() LispValue {
    if len(c) == 0 {
        return Nil
    }
    s := string(c)[0]
    return NewByte(s)
}

func (c ConventionalString) Cdr() LispValue {
    if len(c) < 2 {
        return Nil
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
