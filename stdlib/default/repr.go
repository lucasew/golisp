package stdlib

import (
	"github.com/lucasew/golisp/data"
	"github.com/lucasew/golisp/data/types"
	"github.com/lucasew/golisp/stdlib/default/enforce"
)

func init() {
	register("repr", Repr)
}

func Repr(v ...data.LispValue) (data.LispValue, error) {
	err := enforce.Length(v, 1)
	if err != nil {
		return types.Nil, err
	}
	return types.NewString(v[0].Repr()), nil
}
