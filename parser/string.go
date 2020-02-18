package parser

import (
    "github.com/lucasew/golisp/data"
    "github.com/lucasew/golisp/data/types"
    "github.com/lucasew/golisp/lex"
    "fmt"
)

func ParseString(ctx *lex.Context) (data.LispValue, error) {
    b, ok := ctx.GetByte()
    if !ok {
        return data.Nil, fmt.Errorf("%w: string", ErrEOFWhile)
    }
    if !b.IsStringMark() {
        return data.Nil, fmt.Errorf("%w: string", ErrInvalidEntryPoint)
    }
    begin := ctx.Index() + 1
    for {
        ctx.Increment()
        b, ok := ctx.GetByte()
        if !ok {
            return data.Nil, fmt.Errorf("%w: string", ErrPrematureEOF)
        }
        if b.IsStringMark() {
            s := ctx.Slice(begin, ctx.Index())
            ctx.Increment()
            return types.NewConventionalString(s), nil
        }
    }
}
