package parser

import (
    "github.com/lucasew/golisp/data"
    "github.com/lucasew/golisp/data/types"
    "github.com/lucasew/golisp/lex"
    "errors"
)

func ParseString(ctx *lex.Context) (data.LispValue, error) {
    b, ok := ctx.GetByte()
    if !ok {
        return data.Nil, errors.New("eof when parsing string")
    }
    if !b.IsStringMark() {
        return data.Nil, errors.New("invalid entry point for string")
    }
    begin := ctx.Index() + 1
    for {
        ctx.Increment()
        b, ok := ctx.GetByte()
        if !ok {
            return data.Nil, errors.New("eof when parsing string body")
        }
        if b.IsStringMark() {
            s := ctx.Slice(begin, ctx.Index())
            ctx.Increment()
            return types.NewConventionalString(s), nil
        }
    }
}
