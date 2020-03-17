package stdlib

import (
	"context"
	"github.com/lucasew/golisp/data"
	"github.com/lucasew/golisp/data/types"
	"github.com/lucasew/golisp/utils/enforce"
)

func init() {
	register("print", Print)
	register("println", Println)
}

func Print(ctx context.Context, v ...data.LispValue) (data.LispValue, error) {
	if len(v) == 0 {
		return types.NewString(""), nil
	}
	err := enforce.Validate(enforce.Length(v, 1), enforce.String(v, 1))
	if err != nil {
		return types.Nil, err
	}
	s := v[0].(data.LispString)
	print(s.ToString())
	return s, nil
}

func Println(ctx context.Context, v ...data.LispValue) (data.LispValue, error) {
	if len(v) == 0 {
		return types.NewString(""), nil
	}
	err := enforce.Validate(enforce.Length(v, 1), enforce.String(v, 1))
	if err != nil {
		return types.Nil, err
	}
	s := v[0].(data.LispString)
	println(s.ToString())
	return s, nil
}
