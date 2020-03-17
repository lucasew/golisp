package stdlib

import (
	"context"
	"github.com/lucasew/golisp/data"
	"github.com/lucasew/golisp/data/types"
	"github.com/lucasew/golisp/data/types/iterator"
	"github.com/lucasew/golisp/utils/enforce"
)

func init() {
	register("map", Map)
	register("filter", Filter)
	register("reduce", Reduce)
}

func Map(ctx context.Context, v ...data.LispValue) (data.LispValue, error) {
	err := enforce.Validate(enforce.Length(v, 2), enforce.Function(v, 1))
	if err != nil {
		return types.Nil, err
	}
	fn := v[0].(data.LispFunction)
	lst, err := iterator.NewIterator(v[1])
	if err != nil {
		return types.Nil, err
	}
	return iterator.NewMapIterator(lst, fn), nil
}

func Filter(ctx context.Context, v ...data.LispValue) (data.LispValue, error) {
	err := enforce.Validate(enforce.Length(v, 2), enforce.Function(v, 1))
	if err != nil {
		return types.Nil, err
	}
	fn := v[0].(data.LispFunction)
	lst, err := iterator.NewIterator(v[1])
	if err != nil {
		return types.Nil, err
	}
	return iterator.NewFilterIterator(ctx, lst, fn), nil
}

func Reduce(ctx context.Context, v ...data.LispValue) (data.LispValue, error) {
	err := enforce.Validate(enforce.Length(v, 2), enforce.Function(v, 1))
	if err != nil {
		return types.Nil, err
	}
	fn := v[0].(data.LispFunction)
	lst, err := iterator.NewIterator(v[1])
	if err != nil {
		return types.Nil, err
	}
	ret := lst.Next(ctx)
next:
	select {
	case <-ctx.Done():
		return types.Nil, data.ErrContextCancelled
	default:
		if lst.IsEnd(ctx) {
			return ret, nil
		}
		ret, err = fn.LispCall(ctx, ret, lst.Next(ctx))
		if err != nil {
			return types.Nil, err
		}
	}
	goto next
}
