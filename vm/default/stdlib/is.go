package stdlib

import (
    "github.com/lucasew/golisp/data"
    "github.com/lucasew/golisp/data/convert"
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

func IsNumber(v data.LispValue) (data.LispValue, error) {
    _, is := v.(data.LispNumber)
    return convert.NewLispValue(is)
}

func IsString(v data.LispValue) (data.LispValue, error) {
    _, is := v.(data.LispString)
    return convert.NewLispValue(is)
}

func IsSymbol(v data.LispValue) (data.LispValue, error) {
    _, is := v.(types.Symbol)
    return convert.NewLispValue(is)
}

func IsFunction(v data.LispValue) (data.LispValue, error) {
    _, is := v.(data.LispFunction)
    return convert.NewLispValue(is)
}

func IsFunctionNative(v data.LispValue) (data.LispValue, error) {
    f, is := v.(data.LispFunction)
    if is {
        return convert.NewLispValue(f.IsFunctionNative())
    }
    return convert.NewLispValue(false)
}

func IsAtom(v data.LispValue) (data.LispValue, error) {
    _, is := v.(types.Atom)
    return convert.NewLispValue(is)
}

func IsCons(v data.LispValue) (data.LispValue, error) {
    _, is := v.(data.LispCons)
    return convert.NewLispValue(is)
}

