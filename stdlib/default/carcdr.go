package stdlib

import (
	"github.com/lucasew/golisp/data"
)

func init() {
	register("car", Car)
	register("cdr", Cdr)
}

func Car(v data.LispCons) (data.LispValue, error) {
	return v.Car(), nil
}

func Cdr(v data.LispCons) (data.LispValue, error) {
	return v.Cdr(), nil
}
