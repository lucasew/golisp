package stdlib

import (
	"github.com/lucasew/golisp/data"
	// "github.com/lucasew/golisp/data/convert"
	"github.com/lucasew/golisp/data/types"
	"github.com/lucasew/golisp/data/types/raw"
	"github.com/lucasew/golisp/data/types/test"
	"github.com/lucasew/golisp/utils/enforce"
	"reflect"
)

func init() {
	register("eqref", EqRef)
	register("eqd", EqDeep)
	register("eqt", EqType)
}

func EqRef(v ...data.LispValue) (data.LispValue, error) {
	err := enforce.Validate(enforce.Length(v, 2))
	if err != nil {
		return types.Nil, err
	}
	a := v[0]
	b := v[1]
	return raw.NewLispWrapper(a == b), nil
}

func EqDeep(v ...data.LispValue) (data.LispValue, error) {
	err := enforce.Validate(enforce.Length(v, 2))
	if err != nil {
		return types.Nil, err
	}
	a := v[0]
	b := v[1]
	return raw.NewLispWrapper(reflect.DeepEqual(a, b)), nil
}

func EqType(v ...data.LispValue) (data.LispValue, error) {
	err := enforce.Validate(enforce.Length(v, 2))
	if err != nil {
		return types.Nil, err
	}
	a := v[0]
	b := v[1]
	return raw.NewLispWrapper(test.IsSameType(a, b)), nil
}
