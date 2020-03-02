package stdlib

import (
	"fmt"
	"github.com/lucasew/golisp/data"
	"github.com/lucasew/golisp/data/types"
	"github.com/lucasew/golisp/data/types/iterator"
	"github.com/lucasew/golisp/data/types/number"
	"github.com/lucasew/golisp/stdlib/default/enforce"
)

func init() {
	register("range", Range)
	register("new-iterator", NewIterator)
	register("collect", Collect)
	register("next", Next)
	register("is-end", IsEnd)
}

func Range(v ...data.LispValue) (data.LispValue, error) {
	a := v[0]
	to, ok := a.(number.LispInt)
	if !ok {
		return types.Nil, fmt.Errorf("first argument must be an integer")
	}
	to_num, _ := to.Uint64()
	return iterator.NewRangeIteratorTo(int(to_num)), nil
}

func NewIterator(v ...data.LispValue) (data.LispValue, error) {
	err := enforce.Validate(enforce.Length(v, 1))
	if err != nil {
		return types.Nil, err
	}
	return iterator.NewIterator(v[0])
}

func Collect(v ...data.LispValue) (data.LispValue, error) {
	err := enforce.Validate(enforce.Length(v, 1), enforce.Iterator(v[0], 1))
	if err != nil {
		return types.Nil, err
	}
	iter := v[0].(data.LispIterator)
	ret := []data.LispValue{}
	for !iter.IsEnd() {
		ret = append(ret, iter.Next())
	}
	return types.NewCons(ret...), nil
}

func Next(v ...data.LispValue) (data.LispValue, error) {
	err := enforce.Validate(enforce.Length(v, 1), enforce.Iterator(v[0], 1))
	if err != nil {
		return types.Nil, err
	}
	iter := v[0].(data.LispIterator)
	if iter.IsEnd() {
		return types.Nil, fmt.Errorf("empty iterator")
	}
	return iter.Next(), nil
}

func IsEnd(v ...data.LispValue) (data.LispValue, error) {
	err := enforce.Validate(enforce.Length(v, 1), enforce.Iterator(v[0], 1))
	if err != nil {
		return types.Nil, err
	}
	iter := v[0].(data.LispIterator)
	if iter.IsEnd() {
		return types.T, nil
	} else {
		return types.Nil, nil
	}
}
