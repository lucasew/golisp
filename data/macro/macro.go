package macro

import (
    "github.com/lucasew/golisp/vm"
    "github.com/lucasew/golisp/data"
    "errors"
)

type LispMacro func(vm.LispVM, data.LispValue) (data.LispValue, error)

func ConvertToLispValue(v interface{}) (data.LispValue, error) {
    r, ok := v.(LispMacro)
    if !ok {
        return data.Nil, errors.New("not a macro")
    }
    return r, nil
}

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
