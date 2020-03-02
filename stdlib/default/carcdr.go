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

func Car(v data.LispCons) (data.LispValue, error) {
    err := enforce.Validate(enforce.Length(v, 1), enforce.CarCdr(v.Car(), 1))
    if err != nil {
        return types.Nil, err
    }
	return v.Car().(data.LispCarCdr).Car(), nil
}

func Cdr(v data.LispCons) (data.LispValue, error) {
    err := enforce.Validate(enforce.Length(v, 1), enforce.CarCdr(v.Car(), 1))
    if err != nil {
        return types.Nil, err
    }
	return v.Car().(data.LispCarCdr).Cdr(), nil
}
