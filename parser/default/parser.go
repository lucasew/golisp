package pdefault

import (
    "github.com/lucasew/golisp/datatypes"
    "github.com/lucasew/golisp/parser"
    "github.com/lucasew/golisp/lex"
    // "github.com/davecgh/go-spew/spew"
)

func Parse(s string) (datatypes.LispValue, error) {
    ctx := lex.NewContext([]byte(s))
    list := datatypes.NewCons().(datatypes.Cons)
    err := error(nil)
    for {
        // println("parse")
        ret, err := GlobalState(&ctx)
        // spew.Dump(ret)
        if err !=  nil {
            return datatypes.Nil, err
        }
        if !parser.IsComment(ret) { // Ignore all comments
            list = append(list, ret)
        }
        erreof := ctx.StateWhitespace()
        if erreof != nil {
            break
        }
    }
    // spew.Dump(list)
    return list, err
}
