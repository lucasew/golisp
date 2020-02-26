package types

import (
	"github.com/lucasew/golisp/data"
)

type _nil uint8

func IsNil(v data.LispValue) bool {
	return v.IsNil()
}

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

func (_nil) Next() data.LispValue {
    return Nil
}

func (_nil) IsEnd() bool {
    return true
}

func (_nil) LispCall(data.LispCons) (data.LispValue, error) {
    return Nil, nil
}

func (_nil) IsFunctionNative() bool {
    return true
}
