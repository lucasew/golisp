package stdlib

import (
	"context"
	"fmt"
	"github.com/lucasew/golisp/data"
	"github.com/lucasew/golisp/data/types"
	"github.com/lucasew/golisp/data/types/iterator"
	"github.com/lucasew/golisp/data/types/number"
	"github.com/lucasew/golisp/utils/enforce"
)

func init() {
	register("range", Range)
	register("new-iterator", NewIterator)
	register("collect", Collect)
	register("next", Next)
	register("is-end", IsEnd)
}

func Range(ctx context.Context, v ...data.LispValue) (data.LispValue, error) {
	a := v[0]
	to, ok := a.(number.LispInt)
	if !ok {
		return types.Nil, fmt.Errorf("first argument must be an integer")
	}
	to_num, _ := to.Uint64()
	return iterator.NewRangeIteratorTo(int(to_num)), nil
}

func NewIterator(ctx context.Context, v ...data.LispValue) (data.LispValue, error) {
	err := enforce.Validate(enforce.Length(v, 1))
	if err != nil {
		return types.Nil, err
	}
	return iterator.NewIterator(v[0])
}

func Collect(ctx context.Context, v ...data.LispValue) (data.LispValue, error) {
	err := enforce.Validate(enforce.Length(v, 1), enforce.Entity("lisp_iterator", v, 1))
	if err != nil {
		return types.Nil, err
	}
	iter := v[0].(data.LispIterator)
	ret := []data.LispValue{}
	for !iter.IsEnd(ctx) {
		ret = append(ret, iter.Next(ctx))
	}
	return types.NewCons(ret...), nil
}

func Next(ctx context.Context, v ...data.LispValue) (data.LispValue, error) {
	err := enforce.Validate(enforce.Length(v, 1), enforce.Entity("lisp_iterator", v, 1))
	if err != nil {
		return types.Nil, err
	}
	iter := v[0].(data.LispIterator)
	if iter.IsEnd(ctx) {
		return types.Nil, fmt.Errorf("empty iterator")
	}
	return iter.Next(ctx), nil
}

func IsEnd(ctx context.Context, v ...data.LispValue) (data.LispValue, error) {
	err := enforce.Validate(enforce.Length(v, 1), enforce.Entity("lisp_iterator", v, 1))
	if err != nil {
		return types.Nil, err
	}
	iter := v[0].(data.LispIterator)
	if iter.IsEnd(ctx) {
		return types.T, nil
	} else {
		return types.Nil, nil
	}
}
