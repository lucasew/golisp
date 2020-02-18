package stdlib

import (
    "github.com/lucasew/golisp/data"
    "github.com/lucasew/golisp/data/convert"
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
    a := v.Car()
    b := v.Cdr().Car()
    return callMethod("Sum", a, b)
}

func Sub(v data.LispCons) (data.LispValue, error) {
    a := v.Car()
    b := v.Cdr().Car()
    return callMethod("Sub", a, b)
}

func Neg(v data.LispCons) (data.LispValue, error) {
    a := v.Car()
    b := v.Cdr().Car()
    return callMethod("Neg", a, b)
}

func Mul(v data.LispCons) (data.LispValue, error) {
    a := v.Car()
    b := v.Cdr().Car()
    return callMethod("Mul", a, b)
}

func Div(v data.LispCons) (data.LispValue, error) {
    a := v.Car()
    b := v.Cdr().Car()
    return callMethod("Div", a, b)
}

func Rem(v data.LispCons) (data.LispValue, error) {
    a := v.Car()
    b := v.Cdr().Car()
    return callMethod("Rem", a, b)
}

func Sqrt(v data.LispCons) (data.LispValue, error) {
    a := v.Car()
    if reflect.ValueOf(a).MethodByName("Sqrt").IsValid() {
        return reflect.ValueOf(a).MethodByName("Sqrt").Call([]reflect.Value{})[0].Interface().(data.LispValue), nil
    } else {
        return data.Nil, fmt.Errorf("invalid value for sqrt")
    }
}

func Exp(v data.LispCons) (data.LispValue, error) {
    a := v.Car()
    b := v.Cdr().Car()
    return callMethod("Exp", a, b)
}

func Abs(v data.LispCons) (data.LispValue, error) {
    a := v.Car()
    if reflect.ValueOf(a).MethodByName("Abs").IsValid() {
        return reflect.ValueOf(a).MethodByName("Abs").Call([]reflect.Value{})[0].Interface().(data.LispValue), nil
    } else {
        return data.Nil, fmt.Errorf("invalid value for abs")
    }
}

func IsZero(v data.LispCons) (data.LispValue, error) {
    n := v.Car()
    num, ok := n.(data.LispNumber)
    if !ok {
        return data.Nil, fmt.Errorf("parameter should be a number")
    }
    return convert.NewLispValue(num.IsZero())
}

func IsInt(v data.LispCons) (data.LispValue, error) {
    n, ok := v.Car().(data.LispNumber)
    if !ok {
        return data.Nil, fmt.Errorf("parameter should be a number")
    }
    return convert.NewLispValue(n.IsInt())
}

func callMethod(method string, a data.LispValue, b data.LispValue) (data.LispValue, error) {
    if reflect.TypeOf(a) != reflect.TypeOf(b) {
        return data.Nil, fmt.Errorf("unsupported %s operation between type %T and %T", method, a, b)
    }
    if reflect.ValueOf(a).MethodByName(method).IsValid() {
        rv := reflect.ValueOf(b)
        ret, ok := reflect.ValueOf(a).MethodByName(method).Call([]reflect.Value{rv})[0].Interface().(data.LispValue)
        if !ok {
            return data.Nil, fmt.Errorf("invalid state: method doesnt returns LispValue")
        }
        return ret, nil
    }
    return data.Nil, fmt.Errorf("invalid state: none of the conditions were satisfied")
}
