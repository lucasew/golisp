package stdlib

import (
	"github.com/lucasew/golisp/data"
	"github.com/lucasew/golisp/data/convert"
	"github.com/lucasew/golisp/data/types"
	"github.com/lucasew/golisp/data/types/test"
	"github.com/lucasew/golisp/utils/enforce"
	"reflect"
)

func init() {
	register("eq", Eq)
	register("eqd", EqDeep)
	register("eqt", EqType)
}

func Eq(v ...data.LispValue) (data.LispValue, error) {
	err := enforce.Validate(enforce.Length(v, 2))
	if err != nil {
		return types.Nil, err
	}
	a := v[0]
	b := v[1]
	return convert.NewLispValue(a == b)
}

func EqDeep(v ...data.LispValue) (data.LispValue, error) {
	err := enforce.Validate(enforce.Length(v, 2))
	if err != nil {
		return types.Nil, err
	}
	a := v[0]
	b := v[1]
	return convert.NewLispValue(reflect.DeepEqual(a, b))
}

func EqType(v ...data.LispValue) (data.LispValue, error) {
	err := enforce.Validate(enforce.Length(v, 2))
	if err != nil {
		return types.Nil, err
	}
	a := v[0]
	b := v[1]
	return convert.NewLispValue(test.IsSameType(a, b))
}
