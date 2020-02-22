package stdlib

import (
    "github.com/lucasew/golisp/data"
    "github.com/lucasew/golisp/data/types"
    "github.com/lucasew/golisp/data/convert"
    "github.com/lucasew/golisp/vm/default/stdlib/enforce"
    "reflect"
    "fmt"
)

func init() {
    register("+", Sum)
    register("-", Sub)
    register("neg", Neg)
    register("*", Mul)
    register("/", Div)
    register("%", Rem)
    register("sqrt", Sqrt)
    register("pow", Exp)
    register("abs", Abs)
    register("is-zero", IsZero)
    register("is-int", IsInt)
}

func Sum(v data.LispCons) (data.LispValue, error) {
    return pairOp("Sum", v)
}

func Sub(v data.LispCons) (data.LispValue, error) {
    return pairOp("Sub", v)
}

func Neg(v data.LispCons) (data.LispValue, error) {
    return singleOp("Neg", v)
}

func Mul(v data.LispCons) (data.LispValue, error) {
    return pairOp("Mul", v)
}

func Div(v data.LispCons) (data.LispValue, error) {
    return pairOp("Div", v)
}

func Rem(v data.LispCons) (data.LispValue, error) {
    return pairOp("Rem", v)
}

func Sqrt(v data.LispCons) (data.LispValue, error) {
    return singleOp("Sqrt", v)
}

func Exp(v data.LispCons) (data.LispValue, error) {
    return pairOp("Exp", v)
}

func Abs(v data.LispCons) (data.LispValue, error) {
    return singleOp("Abs", v)
}

func IsZero(v data.LispCons) (data.LispValue, error) {
    n := v.Car()
    num, ok := n.(data.LispNumber)
    if !ok {
        return types.Nil, fmt.Errorf("parameter should be a number")
    }
    return convert.NewLispValue(num.IsZero())
}

func IsInt(v data.LispCons) (data.LispValue, error) {
    n, ok := v.Car().(data.LispNumber)
    if !ok {
        return types.Nil, fmt.Errorf("parameter should be a number")
    }
    return convert.NewLispValue(n.IsInt())
}

func singleOp(method string, v data.LispCons) (data.LispValue, error) {
    a := v.Car()
    err := enforce.Validate(enforce.Length(v, 1), enforce.Number(a))
    if err != nil {
        return types.Nil, err
    }
    if reflect.ValueOf(a).MethodByName(method).IsValid() {
        ret, ok := reflect.ValueOf(a).MethodByName(method).Call([]reflect.Value{})[0].Interface().(data.LispValue)
        if !ok {
            return types.Nil, fmt.Errorf("invalid state: method doesnt returns LispValue")
        }
        return ret, nil
    }
    return types.Nil, fmt.Errorf("invalid state: none of the conditions were satisfied")
}

func pairOp(method string, v data.LispCons) (data.LispValue, error) {
    a := v.Car()
    b := v.Cdr().Car()
    err := enforce.Validate(enforce.Length(v, 2), enforce.SameType(a, b), enforce.Number(a), enforce.Number(b))
    if err != nil {
        return types.Nil, err
    }
    if reflect.ValueOf(a).MethodByName(method).IsValid() {
        rv := reflect.ValueOf(b)
        ret, ok := reflect.ValueOf(a).MethodByName(method).Call([]reflect.Value{rv})[0].Interface().(data.LispValue)
        if !ok {
            return types.Nil, fmt.Errorf("invalid state: method doesnt returns LispValue")
        }
        return ret, nil
    }
    return types.Nil, fmt.Errorf("invalid state: none of the conditions were satisfied")
}
