package stdlib

import (
	"fmt"
	"github.com/lucasew/golisp/data"
	"github.com/lucasew/golisp/data/types"
	"github.com/lucasew/golisp/stdlib/default/enforce"
	"reflect"
)

func init() {
	register(">", MoreThan)
	register(">=", MoreEqualThan)
	register("<", LessThan)
	register("<=", LessEqualThan)
	register("==", Equal)
}

func MoreThan(v ...data.LispValue) (data.LispValue, error) {
	r, err := cmp(v...)
	if err != nil {
		return types.Nil, err
	}
	if r > 0 {
		return types.T, nil
	}
	return types.Nil, nil
}

func MoreEqualThan(v ...data.LispValue) (data.LispValue, error) {
	r, err := cmp(v...)
	if err != nil {
		return types.Nil, err
	}
	if r >= 0 {
		return types.T, nil
	}
	return types.Nil, nil
}

func LessThan(v ...data.LispValue) (data.LispValue, error) {
	r, err := cmp(v...)
	if err != nil {
		return types.Nil, err
	}
	if r < 0 {
		return types.T, nil
	}
	return types.Nil, nil
}

func LessEqualThan(v ...data.LispValue) (data.LispValue, error) {
	r, err := cmp(v...)
	if err != nil {
		return types.Nil, err
	}
	if r <= 0 {
		return types.T, nil
	}
	return types.Nil, nil
}

func Equal(v ...data.LispValue) (data.LispValue, error) {
	r, err := cmp(v...)
	if err != nil {
		return types.Nil, err
	}
	if r == 0 {
		return types.T, nil
	}
	return types.Nil, nil
}

func cmp(v ...data.LispValue) (int, error) {
	err := enforce.Validate(enforce.Length(v, 2), enforce.Number(v, 1), enforce.Number(v, 2), enforce.SameType(v, 1, 2))
	if err != nil {
		return 0, err
	}
	a := v[0]
	b := v[1]
	method := "Cmp"
	if reflect.ValueOf(a).MethodByName(method).IsValid() {
		rv := reflect.ValueOf(b)
		ret, ok := reflect.ValueOf(a).MethodByName(method).Call([]reflect.Value{rv})[0].Interface().(int)
		if !ok {
			return 0, fmt.Errorf("invalid state: method doesnt returns int")
		}
		return ret, nil
	}
	return 0, fmt.Errorf("invalid state: none of the conditions were satisfied")
}
