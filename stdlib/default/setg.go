package stdlib

import (
	"errors"
	"github.com/lucasew/golisp/data"
	"github.com/lucasew/golisp/data/types"
	"github.com/lucasew/golisp/vm"
)

func init() {
	register("setg", Setg)
}

func Setg(env vm.LispVM, v data.LispCons) (data.LispValue, error) {
	key, ok := v.Car().(types.Symbol)
	if !ok {
		return types.Nil, errors.New("invalid output for setg")
	}
	value, err := env.Eval(v.Cdr().Car())
	if err != nil {
		return types.Nil, err
	}
	env.EnvSetGlobal(key.ToString(), value)
	return value, nil
}