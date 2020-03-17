package stdlib

import (
	"context"
	"github.com/lucasew/golisp/data"
	"github.com/lucasew/golisp/data/types"
	"github.com/lucasew/golisp/utils/enforce"
)

func init() {
	register("car", Car)
	register("cdr", Cdr)
}

func Car(ctx context.Context, v ...data.LispValue) (data.LispValue, error) {
	err := enforce.Validate(enforce.Length(v, 1), enforce.Entity("lisp_carcdr", v, 1))
	if err != nil {
		return types.Nil, err
	}
	return v[0].(data.LispCarCdr).Car(), nil
}

func Cdr(ctx context.Context, v ...data.LispValue) (data.LispValue, error) {
	err := enforce.Validate(enforce.Length(v, 1), enforce.Entity("lisp_carcdr", v, 1))
	if err != nil {
		return types.Nil, err
	}
	return v[0].(data.LispCarCdr).Cdr(), nil
}
