package pdefault

import (
    "github.com/lucasew/golisp/data"
    "github.com/lucasew/golisp/lex"
    "github.com/lucasew/golisp/parser"
    "errors"
    "fmt"
)

func GlobalState(ctx *lex.Context) (data.LispValue, error) {
    ctx.StateWhitespace()
    b, ok := ctx.GetByte()
    if !ok {
        return data.Nil, errors.New("eof when parsing global state")
    }
    if b.IsByte('-') {
        ctx.Increment()
        b, ok := ctx.GetByte()
        if !ok {
            return data.Nil, errors.New("eof when parsing comment")
        }
        if b.IsByte('-') {
            ctx.Increment()
            for {
                b, ok := ctx.GetByte()
                if !ok {
                    return data.Nil, errors.New("eof when parsing comment")
                }
                if b.IsByte('\n') {
                    return parser.Comment, nil
                }
                ctx.Increment()
            }
        }
    }
    if b.IsByteColon() {
        return parser.ParseAtom(ctx)
    }
    if b.IsOpenPar() {
        return parser.ParseList(ctx, GlobalState)
    }
    if b.IsStringMark() {
        return parser.ParseString(ctx)
    }
    if b.IsByteNumber() {
        return parser.ParseNumber(ctx)
    }
    if b.IsHash() {
        return ParseSpecialLiteral(ctx)
    }
    if b.IsByteLetter() || b.IsByteSpecialSymbol() {
        return parser.ParseSymbol(ctx)
    }
    // if b.IsClosePar() { // TODO: Test more
    //     // panic("invalid ) token")
    //     return data.Nil, errors.New("invalid ')' token")
    //     // return data.Nil, nil
    // }
    return data.Nil, fmt.Errorf("invalid char: '%s'", string(ctx.MustGetByte()))
}
