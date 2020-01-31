package lex

import (
    "github.com/lucasew/golisp/datatypes"
    "github.com/davecgh/go-spew/spew"
)

func Parse(s string) (datatypes.LispValue, error) {
    ctx := Context{
        data: []byte(s),
        index: 0,
    }
    list := datatypes.NewCons().(datatypes.Cons)
    err := error(nil)
    for {
        println("parse")
        ret, err := ctx.GlobalState()
        spew.Dump(ret)
        if err !=  nil {
            return datatypes.Nil, err
        }
        if !datatypes.IsComment(ret) { // Ignore all comments
            list = append(list, ret)
        }
        erreof := ctx.stateWhitespace()
        if erreof != nil {
            break
        }
    }
    spew.Dump(list)
    return list, err
}
