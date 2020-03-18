package types

import (
	"context"
	"github.com/lucasew/golisp/data"
	"github.com/lucasew/golisp/data/entity"
	"github.com/lucasew/golisp/data/entity/register"
)

func init() {
	register.Register(T.LispEntity())
}

type _t uint8

var T data.LispValue = _t(0)

func (_t) LispEntity() data.LispEntity {
	return entity.Entity{
		"t", func(v data.LispValue) bool {
			return !v.IsNil()
		},
	}
}

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

func (_t) LispCall(context.Context, ...data.LispValue) (data.LispValue, error) {
	return T, nil
}

func (_t) IsFunctionNative() bool {
	return true
}
