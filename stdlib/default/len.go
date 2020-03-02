package stdlib

import (
	"github.com/lucasew/golisp/data"
	"github.com/lucasew/golisp/data/convert"
	"github.com/lucasew/golisp/data/types"
	"github.com/lucasew/golisp/stdlib/default/enforce"
)

func init() {
	register("len", Len)
}

func Len(v ...data.LispValue) (data.LispValue, error) {
	err := enforce.Validate(enforce.Length(v, 1))
	if err != nil {
		return types.Nil, err
	}
	return convert.NewLispValue(_len(v[0]))
}

func _len(v data.LispValue) int {
	l, ok := v.(data.LispLen)
	if ok {
		return l.Len()
	} else {
		if v.IsNil() {
			return 0
		} else {
			return 1
		}
	}
}
