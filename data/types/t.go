package types

import (
	"github.com/lucasew/golisp/data"
)

type _t uint8

var T data.LispValue = _t(0)

func (_t) IsNil() bool {
	return false
}

func (_t) Repr() string {
	return "t"
}

func (_t) LispTypeName() string {
	return "cons"
}

func (_t) Car() data.LispValue {
	return Nil
}

func (_t) Cdr() data.LispCarCdr {
	return Nil
}

func (_t) LispCall(...data.LispValue) (data.LispValue, error) {
	return T, nil
}

func (_t) IsFunctionNative() bool {
	return true
}
