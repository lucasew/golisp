package pdefault

import (
    "github.com/lucasew/golisp/data"
    "github.com/lucasew/golisp/data/types"
    "github.com/lucasew/golisp/lex"
    "github.com/lucasew/golisp/parser"
    "fmt"
)

func GlobalState(ctx *lex.Context) (data.LispValue, error) {
    ctx.StateWhitespace()
    b, ok := ctx.GetByte()
    if !ok {
        return data.Nil, fmt.Errorf("%w: global state", parser.ErrEOFWhile)
    }
    if b.IsByte('-') {
        ctx.Increment()
        b, ok = ctx.GetByte()
        if !ok {
            return data.Nil, fmt.Errorf("%w: comment marker", parser.ErrEOFWhile)
        }
        if b.IsByteNumber() {
            ctx.Decrement()
            return parser.ParseNumber(ctx)
        }
        if b.IsBlank() {
            return types.NewSymbol("-"), nil
        }
        if b.IsByte('-') {
            ctx.Increment()
            for {
                b, ok := ctx.GetByte()
                if !ok {
                    return data.Nil, fmt.Errorf("%w: comment", parser.ErrPrematureEOF)
                }
                if b.IsByte('\n') {
                    return parser.Comment, nil
                }
                ctx.Increment()
            }
        } else {
            sym, err := parser.ParseSymbol(ctx)
            if err != nil {
                return data.Nil, err
            }
            s := fmt.Sprintf("-%s", sym)
            return types.NewSymbol(s), nil
        }
    }
    if b.IsByte('\'') {
        ctx.Increment()
        g, err := GlobalState(ctx)
        return types.NewCons(types.NewSymbol("quote"), g), err
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
    return data.Nil, fmt.Errorf("%w: '%s'", parser.ErrInvalidChar, string(ctx.MustGetByte()))
}
