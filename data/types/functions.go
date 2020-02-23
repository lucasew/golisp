package types

import (
	"github.com/lucasew/golisp/data"
)

type lispFunction func(data.LispCons) (data.LispValue, error)

func (f lispFunction) IsNil() bool {
	return f == nil
}

func (f lispFunction) LispCall(i data.LispCons) (data.LispValue, error) {
	return f(i)
}

func (f lispFunction) Repr() string {
	return "<native function>"
}

func (f lispFunction) IsFunctionNative() bool {
	return true
}

func NewFunction(f func(data.LispCons) (data.LispValue, error)) data.LispFunction {
	return lispFunction(f)
}

func (lispFunction) LispTypeName() string {
	return "function"
}

func IsFunction(v data.LispValue) bool {
	_, ok := v.(data.LispFunction)
	return ok
}

func IsNativeFunction(v data.LispValue) bool {
	fn, ok := v.(data.LispFunction)
	if !ok {
		return false
	}
	return fn.IsFunctionNative()
}
