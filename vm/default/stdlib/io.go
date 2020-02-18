package stdlib

import (
    "github.com/lucasew/golisp/data"
    "github.com/lucasew/golisp/data/convert"
    "errors"
)

func init() {
    register("print", Print)
    register("println", Println)
}

func Print(v data.LispCons) (data.LispValue, error) {
    if v.IsNil() {
        return convert.NewLispValue("")
    }
    s, ok := v.Car().(data.LispString)
    if !ok {
        return data.Nil, errors.New("invalid input")
    }
    print(s.ToString())
    return s, nil
}


func Println(v data.LispCons) (data.LispValue, error) {
    if v.IsNil() {
        return convert.NewLispValue("")
    }
    s, ok := v.Car().(data.LispString)
    if !ok {
        return data.Nil, errors.New("invalid input")
    }
    println(s.ToString())
    return s, nil
}
