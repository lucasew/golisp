package stdlib

import (
	"github.com/lucasew/golisp/data"
	"github.com/lucasew/golisp/data/types"
	"github.com/lucasew/golisp/data/types/number"
	"github.com/lucasew/golisp/data/types/raw"
	"github.com/lucasew/golisp/data/types/test"
	"github.com/lucasew/golisp/utils/enforce"
)

func init() {
	register("is-native", IsNative)
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

func IsNative(v ...data.LispValue) (data.LispValue, error) {
	err := enforce.Validate(enforce.Length(v, 1))
	if err != nil {
		return types.Nil, err
	}
	_, ok := v[0].(raw.LispWrapper)
	return raw.NewLispWrapper(ok), nil
}

func IsNumber(v ...data.LispValue) (data.LispValue, error) {
	err := enforce.Validate(enforce.Length(v, 1))
	if err != nil {
		return types.Nil, err
	}
	return raw.NewLispWrapper(number.IsNumber(v[0])), nil
}

func IsString(v ...data.LispValue) (data.LispValue, error) {
	err := enforce.Validate(enforce.Length(v, 1))
	if err != nil {
		return types.Nil, err
	}
	return raw.NewLispWrapper(types.IsString(v[0])), nil
}

func IsSymbol(v ...data.LispValue) (data.LispValue, error) {
	err := enforce.Validate(enforce.Length(v, 1))
	if err != nil {
		return types.Nil, err
	}
	return raw.NewLispWrapper(types.IsSymbol(v[0])), nil
}

func IsFunction(v ...data.LispValue) (data.LispValue, error) {
	err := enforce.Validate(enforce.Length(v, 1))
	if err != nil {
		return types.Nil, err
	}
	return raw.NewLispWrapper(types.IsFunction(v[0])), nil
}

func IsFunctionNative(v ...data.LispValue) (data.LispValue, error) {
	err := enforce.Validate(enforce.Length(v, 1))
	if err != nil {
		return types.Nil, err
	}
	return raw.NewLispWrapper(types.IsNativeFunction(v[0])), nil
}

func IsAtom(v ...data.LispValue) (data.LispValue, error) {
	err := enforce.Validate(enforce.Length(v, 1))
	if err != nil {
		return types.Nil, err
	}
	return raw.NewLispWrapper(test.IsAtom(v[0])), nil
}

func IsCons(v ...data.LispValue) (data.LispValue, error) {
	err := enforce.Validate(enforce.Length(v, 1))
	if err != nil {
		return types.Nil, err
	}
	return raw.NewLispWrapper(test.IsCons(v[0])), nil
}

func IsMap(v ...data.LispValue) (data.LispValue, error) {
	err := enforce.Validate(enforce.Length(v, 1))
	if err != nil {
		return types.Nil, err
	}
	return raw.NewLispWrapper(test.IsMap(v[0])), nil
}

func IsNamespace(v ...data.LispValue) (data.LispValue, error) {
	err := enforce.Validate(enforce.Length(v, 1))
	if err != nil {
		return types.Nil, err
	}
	return raw.NewLispWrapper(test.IsNamespace(v[0])), nil
}

func IsIterator(v ...data.LispValue) (data.LispValue, error) {
	err := enforce.Validate(enforce.Length(v, 1))
	if err != nil {
		return types.Nil, err
	}
	return raw.NewLispWrapper(test.IsIterator(v[0])), nil
}

func Pass(v ...data.LispValue) (data.LispValue, error) {
	err := enforce.Validate(enforce.Length(v, 1))
	if err != nil {
		return types.Nil, err
	}
	return v[0], nil
}
