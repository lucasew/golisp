package types

import (
    "fmt"
    "github.com/lucasew/golisp/data"
    "testing"
)

type Atom string

func NewAtom(s string) data.LispAtom {
    return Atom(s)
}

func (a Atom) IsNil() bool {
    return string(a) == "nil"
}

func (a Atom) Repr() string {
    return fmt.Sprintf(":%s", a)
}

func (Atom) LispTypeName() string {
    return "atom"
}

func (a Atom) AtomString() string {
    return string(a)
}

func AtomTest(v data.LispValue) func(t *testing.T) {
    return func(t *testing.T) {
        if !IsAtom(v) {
            t.Fail()
        }
    }
}

func IsAtom(v data.LispValue) bool {
    _, ok := v.(data.LispAtom)
    return ok
}
