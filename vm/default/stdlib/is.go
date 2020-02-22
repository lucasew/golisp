package stdlib

import (
    "github.com/lucasew/golisp/data"
    "github.com/lucasew/golisp/data/convert"
    "github.com/lucasew/golisp/vm/default/stdlib/enforce"
    "github.com/lucasew/golisp/data/types"
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
    _, is := v.Car().(data.LispNumber)
    return convert.NewLispValue(is)
}

func IsString(v data.LispCons) (data.LispValue, error) {
    err := enforce.Length(v, 1)
    if err != nil {
        return data.Nil, err
    }
    _, is := v.Car().(data.LispString)
    return convert.NewLispValue(is)
}

func IsSymbol(v data.LispCons) (data.LispValue, error) {
    err := enforce.Length(v, 1)
    if err != nil {
        return data.Nil, err
    }
    _, is := v.Car().(types.Symbol)
    return convert.NewLispValue(is)
}

func IsFunction(v data.LispCons) (data.LispValue, error) {
    err := enforce.Length(v, 1)
    if err != nil {
        return data.Nil, err
    }
    _, is := v.Car().(data.LispFunction)
    return convert.NewLispValue(is)
}

func IsFunctionNative(v data.LispCons) (data.LispValue, error) {
    err := enforce.Length(v, 1)
    if err != nil {
        return data.Nil, err
    }
    f, is := v.Car().(data.LispFunction)
    if is {
        return convert.NewLispValue(f.IsFunctionNative())
    }
    return convert.NewLispValue(false)
}

func IsAtom(v data.LispCons) (data.LispValue, error) {
    err := enforce.Length(v, 1)
    if err != nil {
        return data.Nil, err
    }
    _, is := v.Car().(types.Atom)
    return convert.NewLispValue(is)
}

func IsCons(v data.LispCons) (data.LispValue, error) {
    err := enforce.Length(v, 1)
    if err != nil {
        return data.Nil, err
    }
    _, is := v.Car().(data.LispCons)
    return convert.NewLispValue(is)
}

