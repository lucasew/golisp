package stdlib

import (
	"context"
	"errors"
	"github.com/lucasew/golisp/data"
	eregister "github.com/lucasew/golisp/data/entity/register"
	"github.com/lucasew/golisp/data/types"
	"github.com/lucasew/golisp/data/types/number"
	"github.com/lucasew/golisp/data/types/raw"
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
	register("is-entity", IsEntity)
	register("pass", Pass)
}

func IsEntity(ctx context.Context, v ...data.LispValue) (data.LispValue, error) {
	err := enforce.Validate(
		enforce.Length(v, 2),
		enforce.Entity("lisp_string", v, 1),
	)
	if err != nil {
		return types.Nil, err
	}
	e, ok := eregister.Get(v[0].(data.LispString).ToString())
	if !ok {
		return types.Nil, errors.New("entity not found")
	}
	return raw.NewLispWrapper(e.Isfn(v[1])), nil
}

func IsNative(ctx context.Context, v ...data.LispValue) (data.LispValue, error) {
	err := enforce.Validate(enforce.Length(v, 1))
	if err != nil {
		return types.Nil, err
	}
	_, ok := v[0].(raw.LispWrapper)
	return raw.NewLispWrapper(ok), nil
}

func IsNumber(ctx context.Context, v ...data.LispValue) (data.LispValue, error) {
	err := enforce.Validate(enforce.Length(v, 1))
	if err != nil {
		return types.Nil, err
	}
	return raw.NewLispWrapper(number.IsNumber(v[0])), nil
}

func IsString(ctx context.Context, v ...data.LispValue) (data.LispValue, error) {
	err := enforce.Validate(enforce.Length(v, 1))
	if err != nil {
		return types.Nil, err
	}
	return raw.NewLispWrapper(types.IsString(v[0])), nil
}

func IsSymbol(ctx context.Context, v ...data.LispValue) (data.LispValue, error) {
	err := enforce.Validate(enforce.Length(v, 1))
	if err != nil {
		return types.Nil, err
	}
	return raw.NewLispWrapper(types.IsSymbol(v[0])), nil
}

func IsFunction(ctx context.Context, v ...data.LispValue) (data.LispValue, error) {
	err := enforce.Validate(enforce.Length(v, 1))
	if err != nil {
		return types.Nil, err
	}
	return raw.NewLispWrapper(types.IsFunction(v[0])), nil
}

func IsFunctionNative(ctx context.Context, v ...data.LispValue) (data.LispValue, error) {
	err := enforce.Validate(enforce.Length(v, 1))
	if err != nil {
		return types.Nil, err
	}
	return raw.NewLispWrapper(types.IsNativeFunction(v[0])), nil
}

func IsAtom(ctx context.Context, v ...data.LispValue) (data.LispValue, error) {
	err := enforce.Validate(enforce.Length(v, 1))
	if err != nil {
		return types.Nil, err
	}
	return raw.NewLispWrapper(eregister.Is("lisp_atom", v[0])), nil
}

func IsCons(ctx context.Context, v ...data.LispValue) (data.LispValue, error) {
	err := enforce.Validate(enforce.Length(v, 1))
	if err != nil {
		return types.Nil, err
	}
	return raw.NewLispWrapper(eregister.Is("lisp_cons", v[0])), nil
}

func IsMap(ctx context.Context, v ...data.LispValue) (data.LispValue, error) {
	err := enforce.Validate(enforce.Length(v, 1))
	if err != nil {
		return types.Nil, err
	}
	return raw.NewLispWrapper(eregister.Is("lisp_map", v[0])), nil
}

func IsNamespace(ctx context.Context, v ...data.LispValue) (data.LispValue, error) {
	err := enforce.Validate(enforce.Length(v, 1))
	if err != nil {
		return types.Nil, err
	}
	return raw.NewLispWrapper(eregister.Is("lisp_namespace", v[0])), nil
}

func IsIterator(v ...data.LispValue) (data.LispValue, error) {
	err := enforce.Validate(enforce.Length(v, 1))
	if err != nil {
		return types.Nil, err
	}
	return raw.NewLispWrapper(eregister.Is("lisp_iterator", v[0])), nil
}

func Pass(v ...data.LispValue) (data.LispValue, error) {
	err := enforce.Validate(enforce.Length(v, 1))
	if err != nil {
		return types.Nil, err
	}
	return v[0], nil
}
