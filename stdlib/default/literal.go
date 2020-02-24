package stdlib

import (
	"github.com/lucasew/golisp/data"
	"github.com/lucasew/golisp/data/types"
	"github.com/lucasew/golisp/stdlib/default/enforce"
	"github.com/lucasew/golisp/vm"
)

func init() {
	register("quote", Quote)
	register("list", List)
	register("call", Call)
}

func Quote(env vm.LispVM, v data.LispCons) (data.LispValue, error) {
	err := enforce.Length(v, 1)
	if err != nil {
		return types.Nil, err
	}
	return v.Car(), nil
}

func List(v data.LispCons) (data.LispValue, error) {
	return v, nil
}

func Call(v data.LispCons) (data.LispValue, error) {
	err := enforce.Validate(enforce.Function(v.Car(), 1), enforce.Cons(v.Cdr().Car(), 2))
	if err != nil {
		return types.Nil, err
	}
	fn := v.Car().(data.LispFunction)
	val := v.Cdr().Car().(data.LispCons)
	return fn.LispCall(val)
}
