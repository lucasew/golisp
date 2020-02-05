package parser

import (
    "github.com/lucasew/golisp/datatypes"
    "github.com/lucasew/golisp/lex"
    "errors"
)

func ParseList(ctx *lex.Context, global GlobalStateFunc) (datatypes.LispValue, error) {
    ctx.StateWhitespace()
    b, ok := ctx.GetByte()
    if !ok {
        return datatypes.Nil, errors.New("eof when parsing list")
    }
    if !b.IsOpenPar() {
        return datatypes.Nil, errors.New("invalid entry  point for type list")
    }
    ctx.Increment()
    li := datatypes.NewCons()
    for {
        err := ctx.StateWhitespace()
        if err != nil {
            // return li, nil
            return datatypes.Nil, err
        }
        b, ok = ctx.GetByte()
        if !ok {
            // return li, nil
            return datatypes.Nil, errors.New("eof when parsing list")
        }
        if b.IsClosePar() {
            ctx.Increment()
            return li, nil
        }
        v, err := global(ctx)
        if err != nil {
            return datatypes.Nil, err
        }
        if !IsComment(v) { // Ignore all comments
            li = append(li.(datatypes.Cons), v)
        }
    }
}
