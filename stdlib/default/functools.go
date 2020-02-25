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

func Map(v data.LispCons) (data.LispValue, error) {
	err := enforce.Validate(enforce.Length(v, 2), enforce.Function(v.Car(), 1), enforce.Iterator(v.Cdr().Car(), 2))
	if err != nil {
		return types.Nil, err
	}
	fn := v.Car().(data.LispFunction)
	lst := v.Cdr().Car().(data.LispIterator)
    return iterator.NewMapIterator(lst, fn), nil
}

func Filter(v data.LispCons) (data.LispValue, error) {
	err := enforce.Validate(enforce.Length(v, 2), enforce.Function(v.Car(), 1), enforce.Iterator(v.Cdr().Car(), 2))
	if err != nil {
		return types.Nil, err
	}
	fn := v.Car().(data.LispFunction)
	lst := v.Cdr().Car().(data.LispIterator)
    return iterator.NewFilterIterator(lst, fn), nil
}

func Reduce(v data.LispCons) (data.LispValue, error) {
	err := enforce.Validate(enforce.Length(v, 2), enforce.Function(v.Car(), 1), enforce.Iterator(v.Cdr().Car(), 2))
	if err != nil {
		return types.Nil, err
	}
	fn := v.Car().(data.LispFunction)
	lst := v.Cdr().Car().(data.LispIterator)
	ret := lst.Next()
next:
	if lst.IsEnd() {
		return ret, nil
	}
	ret, err = fn.LispCall(types.NewCons(ret, lst.Next()))
	if err != nil {
		return types.Nil, err
	}
	goto next
}
