package stdlib

import (
	"fmt"
	"github.com/lucasew/golisp/data"
	"github.com/lucasew/golisp/data/types"
	"github.com/lucasew/golisp/data/types/iterator"
	"github.com/lucasew/golisp/data/types/lambda"
	"github.com/lucasew/golisp/stdlib/default/enforce"
	"github.com/lucasew/golisp/vm"
)

func init() {
	register("lambda", Lambda)
}

func Lambda(vm vm.LispVM, v ...data.LispValue) (data.LispValue, error) {
	err := enforce.Cons(v[0], 1)
	if err != nil {
		return types.Nil, err
	}
	old_params := iterator.NewConsIterator(v[0].(data.LispCons))
	params := []string{}
	for !old_params.IsEnd() {
		param, ok := old_params.Next().(types.Symbol)
		if !ok {
			return types.Nil, fmt.Errorf("parameter name need to be a symbol")
		}
		params = append(params, param.ToString())
	}
	ast := types.NewCons(v[1:]...)
	return lambda.NewLambda(vm, ast, params), nil
}