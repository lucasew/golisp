package stdlib

import (
	"github.com/lucasew/golisp/data"
	"github.com/lucasew/golisp/data/types"
	"github.com/lucasew/golisp/data/types/raw"
	"github.com/lucasew/golisp/utils/enforce"
)

func init() {
	register("not", Not)
	register("and", And)
	register("or", Or)
}

func Not(v ...data.LispValue) (data.LispValue, error) {
	err := enforce.Validate(enforce.Length(v, 1))
	if err != nil {
		return types.Nil, err
	}
	return raw.NewLispWrapper(!v[0].IsNil()), nil
}

func And(v ...data.LispValue) (data.LispValue, error) {
	if v[0].IsNil() {
		return types.Nil, nil
	}
	if len(v) == 1 {
		return v[0], nil
	}
	return And(v[1:]...)
}

func Or(v ...data.LispValue) (data.LispValue, error) {
	if len(v) == 0 {
		return types.Nil, nil
	}
	if !v[0].IsNil() {
		return v[0], nil
	}
	return Or(v[1:]...)
}
