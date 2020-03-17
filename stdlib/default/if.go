package stdlib

import (
	"context"
	"github.com/lucasew/golisp/data"
	"github.com/lucasew/golisp/data/types"
	"github.com/lucasew/golisp/utils/enforce"
	"github.com/lucasew/golisp/vm"
)

func init() {
	register("if", If)
}

func If(ctx context.Context, env vm.LispVM, v ...data.LispValue) (data.LispValue, error) {
	err := enforce.Validate(enforce.Length(v, 3))
	if err != nil {
		return types.Nil, err
	}
	cond, err := env.Eval(ctx, v[0])
	if err != nil {
		return types.Nil, err
	}
	if !cond.IsNil() {
		return env.Eval(ctx, v[1])
	} else {
		return env.Eval(ctx, v[2])
	}
}
