package stdlib

import (
	"github.com/lucasew/golisp/data"
	"github.com/lucasew/golisp/data/types"
	"github.com/lucasew/golisp/stdlib/default/enforce"
)

func init() {
	register("print", Print)
	register("println", Println)
}

func Print(v ...data.LispValue) (data.LispValue, error) {
	if len(v) == 0 {
		return types.NewString(""), nil
	}
	err := enforce.Validate(enforce.Length(v, 1), enforce.String(v[0], 1))
	if err != nil {
		return types.Nil, err
	}
	s := v[0].(data.LispString)
	print(s.ToString())
	return s, nil
}

func Println(v ...data.LispValue) (data.LispValue, error) {
	if len(v) == 0 {
		return types.NewString(""), nil
	}
	err := enforce.Validate(enforce.Length(v, 1), enforce.String(v[0], 1))
	if err != nil {
		return types.Nil, err
	}
	s := v[0].(data.LispString)
	println(s.ToString())
	return s, nil
}
