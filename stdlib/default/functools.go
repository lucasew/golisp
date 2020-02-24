package stdlib

import (
	"github.com/lucasew/golisp/data"
	"github.com/lucasew/golisp/data/types"
	"github.com/lucasew/golisp/stdlib/default/enforce"
)

func init() {
	register("map", Map)
	register("reduce", Reduce)
}

func Map(v data.LispCons) (data.LispValue, error) {
	err := enforce.Validate(enforce.Length(v, 2), enforce.Function(v.Car(), 1), enforce.Cons(v.Cdr().Car(), 2))
	if err != nil {
		return types.Nil, err
	}
	fn := v.Car().(data.LispFunction)
	lst := v.Cdr().Car().(data.LispCons)
	ret := make([]data.LispValue, lst.Len())
	for i := 0; i < lst.Len(); i++ {
		v := types.NewCons(lst.Get(i))
		ret[i], err = fn.LispCall(v)
		if err != nil {
			return types.Nil, err
		}
	}
	return types.NewCons(ret...), nil
}

func Reduce(v data.LispCons) (data.LispValue, error) {
	err := enforce.Validate(enforce.Length(v, 2), enforce.Function(v.Car(), 1), enforce.Cons(v.Cdr().Car(), 2))
	if err != nil {
		return types.Nil, err
	}
	fn := v.Car().(data.LispFunction)
	lst := v.Cdr().Car().(data.LispCarCdr)
	ret := lst.Car()
next:
	if lst.Cdr().IsNil() {
		return ret, nil
	}
	lst = lst.Cdr()
	ret, err = fn.LispCall(types.NewCons(ret, lst.Car()))
	if err != nil {
		return types.Nil, err
	}
	goto next
}