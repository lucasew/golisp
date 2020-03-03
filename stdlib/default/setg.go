package stdlib

import (
	"github.com/lucasew/golisp/data"
	"github.com/lucasew/golisp/data/types"
	"github.com/lucasew/golisp/stdlib/default/enforce"
	"github.com/lucasew/golisp/vm"
)

func init() {
	register("setg", Setg)
}

func Setg(env vm.LispVM, v ...data.LispValue) (data.LispValue, error) {
	err := enforce.Validate(enforce.Length(v, 2), enforce.Symbol(v, 1))
	if err != nil {
		return types.Nil, err
	}
	key := v[0].(types.Symbol)
	value, err := env.Eval(v[1])
	if err != nil {
		return types.Nil, err
	}
	env.EnvSetGlobal(key.ToString(), value)
	return value, nil
}
