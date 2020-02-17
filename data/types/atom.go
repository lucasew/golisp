package types

import (
    "fmt"
    "github.com/lucasew/golisp/data"
)

type Atom string

func NewAtom(s string) data.LispValue {
    return Atom(s)
}

func (a Atom) IsNil() bool {
    return string(a) == "nil"
}

func (Atom) Car() data.LispValue {
    return data.Nil
}

func (Atom) Cdr() data.LispValue {
    return data.Nil
}

func (a Atom) Repr() string {
    return fmt.Sprintf(":%s", a)
}
