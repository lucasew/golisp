package types

import (
	"github.com/lucasew/golisp/data"
)

type _nil uint8

var Nil data.LispCarCdr = _nil(0)

func (_nil) IsNil() bool {
	return true
}

func (_nil) Car() data.LispValue {
	return Nil
}

func (_nil) Cdr() data.LispCarCdr {
	return Nil
}

func (_nil) Repr() string {
	return "nil"
}

func (_nil) LispTypeName() string {
	return "cons"
}

func IsNil(v data.LispValue) bool {
	return v.IsNil()
}
