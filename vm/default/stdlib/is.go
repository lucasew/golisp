package stdlib

import (
	"github.com/lucasew/golisp/data"
	"github.com/lucasew/golisp/data/convert"
	"github.com/lucasew/golisp/data/types"
	"github.com/lucasew/golisp/vm/default/stdlib/enforce"
)

func init() {
	register("is-number", IsNumber)
	register("is-string", IsString)
	register("is-symbol", IsSymbol)
	register("is-function", IsFunction)
	register("is-function-native", IsFunctionNative)
	register("is-atom", IsAtom)
	register("is-cons", IsCons)
}

func IsNumber(v data.LispCons) (data.LispValue, error) {
	err := enforce.Length(v, 1)
	if err != nil {
		return types.Nil, err
	}
	return convert.NewLispValue(types.IsNumber(v.Car()))
}

func IsString(v data.LispCons) (data.LispValue, error) {
	err := enforce.Length(v, 1)
	if err != nil {
		return types.Nil, err
	}
	return convert.NewLispValue(types.IsString(v.Car()))
}

func IsSymbol(v data.LispCons) (data.LispValue, error) {
	err := enforce.Length(v, 1)
	if err != nil {
		return types.Nil, err
	}
	return convert.NewLispValue(types.IsSymbol(v.Car()))
}

func IsFunction(v data.LispCons) (data.LispValue, error) {
	err := enforce.Length(v, 1)
	if err != nil {
		return types.Nil, err
	}
	return convert.NewLispValue(types.IsFunction(v.Car()))
}

func IsFunctionNative(v data.LispCons) (data.LispValue, error) {
	err := enforce.Length(v, 1)
	if err != nil {
		return types.Nil, err
	}
	return convert.NewLispValue(types.IsNativeFunction(v.Car()))
}

func IsAtom(v data.LispCons) (data.LispValue, error) {
	err := enforce.Length(v, 1)
	if err != nil {
		return types.Nil, err
	}
	return convert.NewLispValue(types.IsAtom(v.Car()))
}

func IsCons(v data.LispCons) (data.LispValue, error) {
	err := enforce.Length(v, 1)
	if err != nil {
		return types.Nil, err
	}
	return convert.NewLispValue(types.IsCons(v.Car()))
}
