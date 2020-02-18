package parser

import (
    "github.com/lucasew/golisp/data"
    "github.com/lucasew/golisp/data/types"
    "github.com/lucasew/golisp/lex"
    "fmt"
)

func ParseSymbol(ctx *lex.Context) (data.LispValue, error) {
    b, ok := ctx.GetByte()
    if !ok {
        return data.Nil, fmt.Errorf("%w: symbol", ErrEOFWhile)
    }
    if !(b.IsByteLetter() || b.IsByteSpecialSymbol()) {
        return data.Nil, fmt.Errorf("%w: symbol", ErrInvalidEntryPoint)
    }
    begin := ctx.Index()
    for {
        ctx.Increment()
        b, ok = ctx.GetByte()
        if !ok {
            return data.Nil, fmt.Errorf("%w: symbol", ErrPrematureEOF)
        }
        if !(b.IsByteLetter() || b.IsByteSpecialSymbol()) {
            s := ctx.Slice(begin, ctx.Index())
            return types.NewSymbol(s), nil
        }
    }
}
