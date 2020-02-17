package types

import (
    "github.com/lucasew/golisp/data"
)

type lispFunction func(data.LispValue)(data.LispValue, error)

func (f lispFunction) Car() data.LispValue {
    return data.Nil
}

func (f lispFunction) Cdr() data.LispValue {
    return data.Nil
}

func (f lispFunction) IsNil() bool {
    return f == nil
}

func (f lispFunction) LispCall(i data.LispValue) (data.LispValue, error) {
    return f(i)
}

func (f lispFunction) Repr() string {
    return "<native function>"
}

func (f lispFunction) IsFunctionNative() bool {
    return true
}

func NewFunction(f func(data.LispValue)(data.LispValue, error)) data.LispFunction {
    return lispFunction(f)
}
