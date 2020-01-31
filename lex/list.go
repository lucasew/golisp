package lex

import (
    "github.com/lucasew/golisp/datatypes"
    "errors"
)

func (ctx *Context) ParseList() (datatypes.LispValue, error) {
    ctx.stateWhitespace()
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
        err := ctx.stateWhitespace()
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
        v, err := ctx.GlobalState()
        if err != nil {
            return datatypes.Nil, err
        }
        if !datatypes.IsComment(v) { // Ignore all comments
            li = append(li.(datatypes.Cons), v)
        }
    }
}
