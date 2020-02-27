package test

import (
	"github.com/lucasew/golisp/data"
	"reflect"
)

func IsSameType(a, b data.LispValue) bool {
	return reflect.TypeOf(a) == reflect.TypeOf(b)
}
