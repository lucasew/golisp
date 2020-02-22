package pdefault

import (
    "github.com/lucasew/golisp/data"
    "github.com/lucasew/golisp/data/types"
    "github.com/lucasew/golisp/parser"
    "github.com/lucasew/golisp/lex"
    // "github.com/davecgh/go-spew/spew"
)

func Parse(s string) (data.LispValue, error) {
    ctx := lex.NewContext([]byte(s))
    list := types.NewCons().(types.Cons)
    err := error(nil)
    for {
        // println("parse")
        ret, err := GlobalState(&ctx)
        // spew.Dump(ret)
        if err !=  nil {
            return types.Nil, err
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
