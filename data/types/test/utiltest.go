package test

import (
	"github.com/lucasew/golisp/data"
	"testing"
)

func NewTestHelper(f func(data.LispValue) bool) func(v data.LispValue) func(*testing.T) {
	return func(v data.LispValue) func(t *testing.T) {
		return func(t *testing.T) {
			if !f(v) {
				t.Fail()
			}
		}
	}
}

func IsAtom(v data.LispValue) bool {
	_, ok := v.(data.LispAtom)
	return ok
}

func IsCarCdr(v data.LispValue) bool {
	_, ok := v.(data.LispCarCdr)
	return ok
}

func IsCons(v data.LispValue) bool {
	_, ok := v.(data.LispCons)
	return ok
}

func IsFunction(v data.LispValue) bool {
	_, ok := v.(data.LispFunction)
	return ok
}

func IsIterator(v data.LispValue) bool {
	_, ok := v.(data.LispIterator)
	return ok
}

func IsLen(v data.LispValue) bool {
	_, ok := v.(data.LispLen)
	return ok
}

func IsMap(v data.LispValue) bool {
	_, ok := v.(data.LispMap)
	return ok
}

func IsNamespace(v data.LispValue) bool {
	_, ok := v.(data.LispNamespace)
	return ok
}

func IsNumber(v data.LispValue) bool {
	_, ok := v.(data.LispNumber)
	return ok
}

func IsPortal(v data.LispValue) bool {
	_, ok := v.(data.LispPortal)
	return ok
}

func IsString(v data.LispValue) bool {
	_, ok := v.(data.LispString)
	return ok
}

func IsValue(v data.LispValue) bool {
	_, ok := v.(data.LispValue)
	return ok
}
