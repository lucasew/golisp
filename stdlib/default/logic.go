package stdlib

import (
	"context"
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

func Not(ctx context.Context, v ...data.LispValue) (data.LispValue, error) {
	err := enforce.Validate(enforce.Length(v, 1))
	if err != nil {
		return types.Nil, err
	}
	return raw.NewLispWrapper(!v[0].IsNil()), nil
}

func And(ctx context.Context, v ...data.LispValue) (data.LispValue, error) {
	if v[0].IsNil() {
		return types.Nil, nil
	}
	if len(v) == 1 {
		return v[0], nil
	}
	select {
	case _ = <-ctx.Done():
		return types.Nil, data.ErrContextCancelled
	default:
		return And(ctx, v[1:]...)
	}
}

func Or(ctx context.Context, v ...data.LispValue) (data.LispValue, error) {
	if len(v) == 0 {
		return types.Nil, nil
	}
	if !v[0].IsNil() {
		return v[0], nil
	}
	select {
	case _ = <-ctx.Done():
		return types.Nil, data.ErrContextCancelled
	default:
		return Or(ctx, v[1:]...)
	}
}
