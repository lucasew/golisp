package lex

import (
    "github.com/lucasew/golisp/datatypes"
    "errors"
    "fmt"
)

func (ctx *Context) GlobalState() (datatypes.LispValue, error) {
    ctx.stateWhitespace()
    b, ok := ctx.GetByte()
    if !ok {
        return datatypes.Nil, errors.New("eof when parsing global state")
    }
    if b.IsByte('-') {
        ctx.Increment()
        b, ok := ctx.GetByte()
        if !ok {
            return datatypes.Nil, errors.New("eof when parsing comment")
        }
        if b.IsByte('-') {
            ctx.Increment()
            for {
                b, ok := ctx.GetByte()
                if !ok {
                    return datatypes.Nil, errors.New("eof when parsing comment")
                }
                if b.IsByte('\n') {
                    return datatypes.Comment, nil
                }
                ctx.Increment()
            }
        }
    }
    if b.IsByteColon() {
        return ctx.ParseAtom()
    }
    if b.IsOpenPar() {
        return ctx.ParseList()
    }
    if b.IsStringMark() {
        return ctx.ParseString()
    }
    if b.IsByteNumber() {
        return ctx.ParseNumber()
    }
    if b.IsHash() {
        return ctx.ParseSpecialLiteral()
    }
    if b.IsByteLetter() || b.IsByteSpecialSymbol() {
        return ctx.ParseSymbol()
    }
    // if b.IsClosePar() { // TODO: Test more
    //     // panic("invalid ) token")
    //     return datatypes.Nil, errors.New("invalid ')' token")
    //     // return datatypes.Nil, nil
    // }
    return datatypes.Nil, fmt.Errorf("invalid char: '%s'", string(ctx.MustGetByte()))
}
