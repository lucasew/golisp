package stdlib

import (
    "github.com/lucasew/golisp/data"
)

func init() {
    register("car", Car)
    register("cdr", Cdr)
}

func Car(v data.LispValue) (data.LispValue, error) {
    return v.Car(), nil
}

func Cdr(v data.LispValue) (data.LispValue, error) {
    return v.Cdr(), nil
}
