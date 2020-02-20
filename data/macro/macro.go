package macro

import (
    "github.com/lucasew/golisp/vm"
    "github.com/lucasew/golisp/data"
)

type LispMacro func(vm.LispVM, data.LispCons) (data.LispValue, error)

func NewLispMacro(f func(vm.LispVM, data.LispCons) (data.LispValue, error)) LispMacro {
    return f
}

func (LispMacro) IsNil() bool {
    return false
}

func (LispMacro) Repr() string {
    return "<native macro>"
}

func (m LispMacro) LispCallMacro(v vm.LispVM, val data.LispCons) (data.LispValue, error) {
    return m(v, val)
}
