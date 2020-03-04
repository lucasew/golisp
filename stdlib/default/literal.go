package stdlib

import (
	"github.com/lucasew/golisp/data"
	"github.com/lucasew/golisp/data/types"
	"github.com/lucasew/golisp/utils/enforce"
	"github.com/lucasew/golisp/vm"
)

func init() {
	register("quote", Quote)
	register("list", List)
	register("call", Call)
}

func Quote(env vm.LispVM, v ...data.LispValue) (data.LispValue, error) {
	err := enforce.Validate(enforce.Length(v, 1))
	if err != nil {
		return types.Nil, err
	}
	return v[0], nil
}

func List(v ...data.LispValue) (data.LispValue, error) {
	return v[0], nil
}

func Call(v ...data.LispValue) (data.LispValue, error) {
	err := enforce.Validate(enforce.Length(v, 2), enforce.Function(v, 1), enforce.Cons(v, 2))
	if err != nil {
		return types.Nil, err
	}
	fn := v[0].(data.LispFunction)
	val := v[1].(data.LispCons)
	return fn.LispCall(val)
}
