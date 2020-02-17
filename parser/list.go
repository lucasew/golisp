package parser

import (
    "github.com/lucasew/golisp/data"
    "github.com/lucasew/golisp/data/types"
    "github.com/lucasew/golisp/lex"
    "errors"
)

func ParseList(ctx *lex.Context, global GlobalStateFunc) (data.LispValue, error) {
    ctx.StateWhitespace()
    b, ok := ctx.GetByte()
    if !ok {
        return data.Nil, errors.New("eof when parsing list")
    }
    if !b.IsOpenPar() {
        return data.Nil, errors.New("invalid entry  point for type list")
    }
    ctx.Increment()
    li := types.NewCons()
    for {
        err := ctx.StateWhitespace()
        if err != nil {
            // return li, nil
            return data.Nil, err
        }
        b, ok = ctx.GetByte()
        if !ok {
            // return li, nil
            return data.Nil, errors.New("eof when parsing list")
        }
        if b.IsClosePar() {
            ctx.Increment()
            return li, nil
        }
        v, err := global(ctx)
        if err != nil {
            return data.Nil, err
        }
        if !IsComment(v) { // Ignore all comments
            li = append(li.(types.Cons), v)
        }
    }
}
