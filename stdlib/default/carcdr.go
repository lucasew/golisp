package stdlib

import (
	"github.com/lucasew/golisp/data"
	"github.com/lucasew/golisp/data/types"
	"github.com/lucasew/golisp/stdlib/default/enforce"
)

func init() {
	register("car", Car)
	register("cdr", Cdr)
}

func Car(v ...data.LispValue) (data.LispValue, error) {
	err := enforce.Validate(enforce.Length(v, 1), enforce.CarCdr(v, 1))
	if err != nil {
		return types.Nil, err
	}
	return v[0].(data.LispCarCdr).Car(), nil
}

func Cdr(v ...data.LispValue) (data.LispValue, error) {
	err := enforce.Validate(enforce.Length(v, 1), enforce.CarCdr(v, 1))
	if err != nil {
		return types.Nil, err
	}
	return v[0].(data.LispCarCdr).Cdr(), nil
}
