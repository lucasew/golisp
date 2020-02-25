package stdlib

import (
	"github.com/lucasew/golisp/data"
	"github.com/lucasew/golisp/data/types"
	"github.com/lucasew/golisp/data/types/iterator"
	"github.com/lucasew/golisp/data/types/number"
	"github.com/lucasew/golisp/stdlib/default/enforce"
	"fmt"
)

func init() {
	register("range", Range)
	register("into-iterator", IntoIterator)
	register("collect", Collect)
	register("next", Next)
}

func Range(v data.LispCons) (data.LispValue, error) {
	a := v.Car()
	to, ok := a.(number.LispInt)
	if !ok {
		return types.Nil, fmt.Errorf("first argument must be an integer")
	}
	to_num, _ := to.Uint64()
	return iterator.NewRangeIteratorTo(int(to_num)), nil
}

func IntoIterator(v data.LispCons) (data.LispValue, error) {
	err := enforce.Validate(enforce.Length(v, 1))
	if err != nil {
		return types.Nil, err
	}
	in := v.Car().(data.LispCons)
    return iterator.NewIterator(in)
}

func Collect(v data.LispCons) (data.LispValue, error) {
	err := enforce.Validate(enforce.Length(v, 1), enforce.Iterator(v.Car(), 1))
	if err != nil {
		return types.Nil, err
	}
	iter := v.Car().(data.LispIterator)
	ret := []data.LispValue{}
	for !iter.IsEnd() {
		ret = append(ret, iter.Next())
	}
	return types.NewCons(ret...), nil
}

func Next(v data.LispCons) (data.LispValue, error) {
	err := enforce.Validate(enforce.Length(v, 1), enforce.Iterator(v.Car(), 1))
	if err != nil {
		return types.Nil, err
	}
	iter := v.Car().(data.LispIterator)
	if iter.IsEnd() {
		return types.Nil, fmt.Errorf("empty iterator")
	}
	return iter.Next(), nil
}
