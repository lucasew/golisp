package stdlib

import (
	"github.com/lucasew/golisp/data"
	"github.com/lucasew/golisp/data/convert"
	"github.com/lucasew/golisp/data/types"
	"github.com/lucasew/golisp/data/types/number"
	"github.com/lucasew/golisp/data/types/test"
	"github.com/lucasew/golisp/stdlib/default/enforce"
)

func init() {
	register("is-number", IsNumber)
	register("is-string", IsString)
	register("is-symbol", IsSymbol)
	register("is-function", IsFunction)
	register("is-function-native", IsFunctionNative)
	register("is-atom", IsAtom)
	register("is-cons", IsCons)
	register("is-map", IsMap)
	register("is-namespace", IsNamespace)
	register("is-iterator", IsIterator)
	register("pass", Pass)
}

func IsNumber(v ...data.LispValue) (data.LispValue, error) {
	err := enforce.Length(v, 1)
	if err != nil {
		return types.Nil, err
	}
	return convert.NewLispValue(number.IsNumber(v[0]))
}

func IsString(v ...data.LispValue) (data.LispValue, error) {
	err := enforce.Length(v, 1)
	if err != nil {
		return types.Nil, err
	}
	return convert.NewLispValue(types.IsString(v[0]))
}

func IsSymbol(v ...data.LispValue) (data.LispValue, error) {
	err := enforce.Length(v, 1)
	if err != nil {
		return types.Nil, err
	}
	return convert.NewLispValue(types.IsSymbol(v[0]))
}

func IsFunction(v ...data.LispValue) (data.LispValue, error) {
	err := enforce.Length(v, 1)
	if err != nil {
		return types.Nil, err
	}
	return convert.NewLispValue(types.IsFunction(v[0]))
}

func IsFunctionNative(v ...data.LispValue) (data.LispValue, error) {
	err := enforce.Length(v, 1)
	if err != nil {
		return types.Nil, err
	}
	return convert.NewLispValue(types.IsNativeFunction(v[0]))
}

func IsAtom(v ...data.LispValue) (data.LispValue, error) {
	err := enforce.Length(v, 1)
	if err != nil {
		return types.Nil, err
	}
	return convert.NewLispValue(test.IsAtom(v[0]))
}

func IsCons(v ...data.LispValue) (data.LispValue, error) {
	err := enforce.Length(v, 1)
	if err != nil {
		return types.Nil, err
	}
	return convert.NewLispValue(test.IsCons(v[0]))
}

func IsMap(v ...data.LispValue) (data.LispValue, error) {
	err := enforce.Length(v, 1)
	if err != nil {
		return types.Nil, err
	}
	return convert.NewLispValue(test.IsMap(v[0]))
}

func IsNamespace(v ...data.LispValue) (data.LispValue, error) {
	err := enforce.Length(v, 1)
	if err != nil {
		return types.Nil, err
	}
	return convert.NewLispValue(test.IsNamespace(v[0]))
}

func IsIterator(v ...data.LispValue) (data.LispValue, error) {
	err := enforce.Length(v, 1)
	if err != nil {
		return types.Nil, err
	}
	return convert.NewLispValue(test.IsIterator(v[0]))
}

func Pass(v ...data.LispValue) (data.LispValue, error) {
	err := enforce.Length(v, 1)
	if err != nil {
		return types.Nil, err
	}
	return v[0], nil
}
