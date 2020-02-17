package parser

import (
    "github.com/lucasew/golisp/data"
    "github.com/lucasew/golisp/data/types"
    "github.com/lucasew/golisp/lex"
    "errors"
)

func ParseSymbol(ctx *lex.Context) (data.LispValue, error) {
    b, ok := ctx.GetByte()
    if !ok {
        return data.Nil, errors.New("eof when parsing symbol")
    }
    if !(b.IsByteLetter() || b.IsByteSpecialSymbol()) {
        return data.Nil, errors.New("invalid entry point when parsing symbol")
    }
    begin := ctx.Index()
    for {
        ctx.Increment()
        b, ok = ctx.GetByte()
        if !ok {
            return data.Nil, errors.New("eof when parsing symbol body")
        }
        if !(b.IsByteLetter() || b.IsByteSpecialSymbol()) {
            s := ctx.Slice(begin, ctx.Index())
            return types.NewSymbol(s), nil
        }
    }
}
