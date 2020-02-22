package stdlib

import (
    "github.com/lucasew/golisp/data"
    "github.com/lucasew/golisp/data/types"
    "github.com/lucasew/golisp/data/convert"
    "reflect"
    "github.com/lucasew/golisp/vm/default/stdlib/enforce"
)

func init() {
    register("eq", Eq)
    register("eqd", EqDeep)
    register("eqt", EqType)
}

func Eq(v data.LispCons) (data.LispValue, error) {
    err := enforce.Length(v, 2)
    if err != nil {
        return types.Nil, err
    }
    a := v.Car()
    b := v.Cdr().Car()
    return convert.NewLispValue(a == b)
}

func EqDeep(v data.LispCons) (data.LispValue, error) {
    err := enforce.Length(v, 2)
    if err != nil {
        return types.Nil, err
    }
    a := v.Car()
    b := v.Cdr().Car()
    return convert.NewLispValue(reflect.DeepEqual(a, b))
}

func EqType(v data.LispCons) (data.LispValue, error) {
    err := enforce.Length(v, 2)
    if err != nil {
        return types.Nil, err
    }
    a := v.Car()
    b := v.Cdr().Car()
    return convert.NewLispValue(types.IsSameType(a, b))
}

