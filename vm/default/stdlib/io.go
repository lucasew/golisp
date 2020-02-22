package stdlib

import (
    "github.com/lucasew/golisp/data"
    "github.com/lucasew/golisp/data/types"
    "github.com/lucasew/golisp/data/convert"
    "github.com/lucasew/golisp/vm/default/stdlib/enforce"
)

func init() {
    register("print", Print)
    register("println", Println)
}

func Print(v data.LispCons) (data.LispValue, error) {
    if v.IsNil() {
        return convert.NewLispValue("")
    }
    err := enforce.Validate(enforce.String(v.Car(), 1), enforce.Length(v, 1))
    if err != nil {
        return types.Nil, err
    }
    s := v.Car().(data.LispString)
    print(s.ToString())
    return s, nil
}


func Println(v data.LispCons) (data.LispValue, error) {
    if v.IsNil() {
        return convert.NewLispValue("")
    }
    err := enforce.Validate(enforce.String(v.Car(), 1), enforce.Length(v, 1))
    if err != nil {
        return types.Nil, err
    }
    s := v.Car().(data.LispString)
    println(s.ToString())
    return s, nil
}
