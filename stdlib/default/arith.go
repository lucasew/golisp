package stdlib

import (
	"fmt"
	"github.com/lucasew/golisp/data"
	"github.com/lucasew/golisp/data/types"
	"github.com/lucasew/golisp/data/types/raw"
	"github.com/lucasew/golisp/utils/enforce"
	"reflect"
)

func init() {
	register("+", Sum)
	register("-", Sub)
	register("neg", Neg)
	register("*", Mul)
	register("/", Div)
	register("%", Rem)
	register("sqrt", Sqrt)
	register("pow", Exp)
	register("abs", Abs)
	register("is-zero", IsZero)
	register("is-int", IsInt)
}

func Sum(v ...data.LispValue) (data.LispValue, error) {
	return pairOp("Sum", v)
}

func Sub(v ...data.LispValue) (data.LispValue, error) {
	return pairOp("Sub", v)
}

func Neg(v ...data.LispValue) (data.LispValue, error) {
	return singleOp("Neg", v)
}

func Mul(v ...data.LispValue) (data.LispValue, error) {
	return pairOp("Mul", v)
}

func Div(v ...data.LispValue) (data.LispValue, error) {
	return pairOp("Div", v)
}

func Rem(v ...data.LispValue) (data.LispValue, error) {
	return pairOp("Rem", v)
}

func Sqrt(v ...data.LispValue) (data.LispValue, error) {
	return singleOp("Sqrt", v)
}

func Exp(v ...data.LispValue) (data.LispValue, error) {
	return pairOp("Exp", v)
}

func Abs(v ...data.LispValue) (data.LispValue, error) {
	return singleOp("Abs", v)
}

func IsZero(v ...data.LispValue) (data.LispValue, error) {
	err := enforce.Validate(enforce.Length(v, 1), enforce.Number(v, 1))
	if err != nil {
		return types.Nil, err
	}
	num := v[0].(data.LispNumber)
	return raw.NewLispWrapper(num.IsZero()), nil
}

func IsInt(v ...data.LispValue) (data.LispValue, error) {
	err := enforce.Validate(enforce.Length(v, 1), enforce.Number(v, 1))
	if err != nil {
		return types.Nil, err
	}
	n := v[0].(data.LispNumber)
	return raw.NewLispWrapper(n.IsInt()), nil
}

func singleOp(method string, v []data.LispValue) (data.LispValue, error) {
	err := enforce.Validate(enforce.Length(v, 1), enforce.Number(v, 1))
	if err != nil {
		return types.Nil, err
	}
	a := v[0]
	if reflect.ValueOf(a).MethodByName(method).IsValid() {
		ret, ok := reflect.ValueOf(a).MethodByName(method).Call([]reflect.Value{})[0].Interface().(data.LispValue)
		if !ok {
			return types.Nil, fmt.Errorf("invalid state: method doesnt returns LispValue")
		}
		return ret, nil
	}
	return types.Nil, fmt.Errorf("invalid state: none of the conditions were satisfied")
}

func pairOp(method string, v []data.LispValue) (data.LispValue, error) {
	err := enforce.Validate(enforce.Length(v, 2), enforce.Number(v, 1), enforce.Number(v, 2), enforce.SameType(v, 1, 2))
	if err != nil {
		return types.Nil, err
	}
	a := v[0]
	b := v[1]
	if reflect.ValueOf(a).MethodByName(method).IsValid() {
		rv := reflect.ValueOf(b)
		ret, ok := reflect.ValueOf(a).MethodByName(method).Call([]reflect.Value{rv})[0].Interface().(data.LispValue)
		if !ok {
			return types.Nil, fmt.Errorf("invalid state: method doesnt returns LispValue")
		}
		return ret, nil
	}
	return types.Nil, fmt.Errorf("invalid state: none of the conditions were satisfied")
}
