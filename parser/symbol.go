package parser

import (
    "github.com/lucasew/golisp/datatypes"
    "github.com/lucasew/golisp/lex"
    "errors"
    // "fmt"
)

func ParseSymbol(ctx *lex.Context) (datatypes.LispValue, error) {
    b, ok := ctx.GetByte()
    if !ok {
        return datatypes.Nil, errors.New("eof when parsing symbol")
    }
    if !(b.IsByteLetter() || b.IsByteSpecialSymbol()) {
        return datatypes.Nil, errors.New("invalid entry point when parsing symbol")
    }
    begin := ctx.Index()
    for {
        ctx.Increment()
        b, ok = ctx.GetByte()
        if !ok {
            return datatypes.Nil, errors.New("eof when parsing symbol body")
        }
        if !(b.IsByteLetter() || b.IsByteSpecialSymbol()) {
            s := ctx.Slice(begin, ctx.Index())
            return datatypes.NewSymbol(s), nil
        }
    }
}
