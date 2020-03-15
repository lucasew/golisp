package types

import (
	"github.com/lucasew/golisp/data"
	"github.com/lucasew/golisp/data/entity/register"
)

func init() {
	register.Register("native_function", func(v data.LispValue) bool {
		_, ok := v.(NativeFunction)
		return ok
	})
}

type NativeFunction func(...data.LispValue) (data.LispValue, error)

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

func (f NativeFunction) IsNil() bool {
	return f == nil
}

func (f NativeFunction) LispCall(i ...data.LispValue) (data.LispValue, error) {
	return f(i...)
}

func (f NativeFunction) Repr() string {
	return "<native function>"
}

func (f NativeFunction) IsFunctionNative() bool {
	return true
}

func NewFunction(f func(...data.LispValue) (data.LispValue, error)) data.LispFunction {
	return NativeFunction(f)
}

func (NativeFunction) LispTypeName() string {
	return "function"
}
