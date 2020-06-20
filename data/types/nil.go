package types

import (
	"context"
	"fmt"
	"github.com/lucasew/golisp/data"
	"github.com/lucasew/golisp/data/entity"
	"github.com/lucasew/golisp/data/entity/register"
)

func init() {
	register.Register(Nil.LispEntity())
}

type _nil uint8

func IsNil(v data.LispValue) bool {
	return v.IsNil()
}

var Nil data.LispCarCdr = _nil(0)

func (_nil) LispEntity() data.LispEntity {
	return entity.Entity{
		"nil", func(v data.LispValue) bool {
			return v.IsNil()
		},
	}
}

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

func (_nil) Next(context.Context) data.LispValue {
	return Nil
}

func (_nil) IsEnd(context.Context) bool {
	return true
}

func (_nil) LispCall(context.Context, ...data.LispValue) (data.LispValue, error) {
	return Nil, fmt.Errorf("you are calling nil")
}

func (_nil) IsFunctionNative() bool {
	return true
}
