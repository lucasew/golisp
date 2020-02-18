package macro

import (
    "github.com/lucasew/golisp/vm"
    "github.com/lucasew/golisp/data"
)

type LispMacro func(vm.LispVM, data.LispValue) (data.LispValue, error)

func NewLispMacro(f func(vm.LispVM, data.LispValue) (data.LispValue, error)) LispMacro {
    return f
}

func (LispMacro) Car() data.LispValue {
    return data.Nil
}

func (LispMacro) Cdr() data.LispValue {
    return data.Nil
}

func (LispMacro) IsNil() bool {
    return false
}

func (LispMacro) Repr() string {
    return "<native macro>"
}

func (m LispMacro) LispCallMacro(v vm.LispVM, val data.LispValue) (data.LispValue, error) {
    return m(v, val)
}