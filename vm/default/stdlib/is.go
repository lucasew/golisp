package stdlib

import (
    "github.com/lucasew/golisp/data"
    "github.com/lucasew/golisp/data/is"
    "github.com/lucasew/golisp/data/convert"
    "github.com/lucasew/golisp/vm/default/stdlib/enforce"
)

func init() {
    register("is-number", IsNumber)
    register("is-string", IsString)
    register("is-symbol", IsSymbol)
    register("is-function", IsFunction)
    register("is-function-native", IsFunctionNative)
    register("is-atom", IsAtom)
    register("is-cons", IsCons)
}

func IsNumber(v data.LispCons) (data.LispValue, error) {
    err := enforce.Length(v, 1)
    if err != nil {
        return data.Nil, err
    }
    return convert.NewLispValue(is.Number(v.Car()))
}

func IsString(v data.LispCons) (data.LispValue, error) {
    err := enforce.Length(v, 1)
    if err != nil {
        return data.Nil, err
    }
    return convert.NewLispValue(is.String(v.Car()))
}

func IsSymbol(v data.LispCons) (data.LispValue, error) {
    err := enforce.Length(v, 1)
    if err != nil {
        return data.Nil, err
    }
    return convert.NewLispValue(is.Symbol(v.Car()))
}

func IsFunction(v data.LispCons) (data.LispValue, error) {
    err := enforce.Length(v, 1)
    if err != nil {
        return data.Nil, err
    }
    return convert.NewLispValue(is.Function(v.Car()))
}

func IsFunctionNative(v data.LispCons) (data.LispValue, error) {
    err := enforce.Length(v, 1)
    if err != nil {
        return data.Nil, err
    }
    return convert.NewLispValue(is.NativeFunction(v.Car()))
}

func IsAtom(v data.LispCons) (data.LispValue, error) {
    err := enforce.Length(v, 1)
    if err != nil {
        return data.Nil, err
    }
    return convert.NewLispValue(is.Atom(v.Car()))
}

func IsCons(v data.LispCons) (data.LispValue, error) {
    err := enforce.Length(v, 1)
    if err != nil {
        return data.Nil, err
    }
    return convert.NewLispValue(is.Cons(v.Car()))
}

