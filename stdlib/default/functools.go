package stdlib

import (
	"github.com/lucasew/golisp/data"
	"github.com/lucasew/golisp/data/types"
	"github.com/lucasew/golisp/data/types/iterator"
	"github.com/lucasew/golisp/stdlib/default/enforce"
)

func init() {
	register("map", Map)
	register("filter", Filter)
	register("reduce", Reduce)
}

func Map(v ...data.LispValue) (data.LispValue, error) {
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

func Filter(v ...data.LispValue) (data.LispValue, error) {
	err := enforce.Validate(enforce.Length(v, 2), enforce.Function(v, 1))
	if err != nil {
		return types.Nil, err
	}
	fn := v[0].(data.LispFunction)
	lst, err := iterator.NewIterator(v[1])
	if err != nil {
		return types.Nil, err
	}
	return iterator.NewFilterIterator(lst, fn), nil
}

func Reduce(v ...data.LispValue) (data.LispValue, error) {
	err := enforce.Validate(enforce.Length(v, 2), enforce.Function(v, 1))
	if err != nil {
		return types.Nil, err
	}
	fn := v[0].(data.LispFunction)
	lst, err := iterator.NewIterator(v[1])
	if err != nil {
		return types.Nil, err
	}
	ret := lst.Next()
next:
	if lst.IsEnd() {
		return ret, nil
	}
	ret, err = fn.LispCall(ret, lst.Next())
	if err != nil {
		return types.Nil, err
	}
	goto next
}
