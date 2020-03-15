package lispfunc

import (
	"github.com/lucasew/golisp/data"
	"github.com/lucasew/golisp/data/entity"
	"github.com/lucasew/golisp/data/types"
	"github.com/lucasew/golisp/utils/enforce"
)

func EntityAsLispFunc(e entity.Entity) func(...data.LispValue) (data.LispValue, error) {
	return func(v ...data.LispValue) (data.LispValue, error) {
		err := enforce.Length(v, 1)()
		if err != nil {
			return types.Nil, err
		}
		if e.Isfn(v[0]) {
			return types.T, nil
		}
		return types.Nil, nil
	}
}
