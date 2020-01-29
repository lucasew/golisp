package datatypes

import (
    "fmt"
)

type Atom string

func NewAtom(s string) LispValue {
    return Atom(s)
}

func (a Atom) IsNil() bool {
    return string(a) == "nil"
}

func (Atom) Car() LispValue {
    return Nil
}

func (Atom) Cdr() LispValue {
    return Nil
}

func (a Atom) Repr() string {
    return fmt.Sprintf(":%s", a)
}
