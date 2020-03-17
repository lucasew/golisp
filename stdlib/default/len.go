package stdlib

import (
	"context"
	"github.com/lucasew/golisp/data"
	"github.com/lucasew/golisp/data/types"
	"github.com/lucasew/golisp/data/types/raw"
	"github.com/lucasew/golisp/utils/enforce"
)

func init() {
	register("len", Len)
}

func Len(ctx context.Context, v ...data.LispValue) (data.LispValue, error) {
	err := enforce.Validate(enforce.Length(v, 1))
	if err != nil {
		return types.Nil, err
	}
	return raw.NewLispWrapper(_len(v[0])), nil
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
