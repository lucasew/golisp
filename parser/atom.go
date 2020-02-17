package parser

import (
    "github.com/lucasew/golisp/data"
    "github.com/lucasew/golisp/data/types"
    "github.com/lucasew/golisp/lex"
    "errors"
)

func ParseAtom(ctx *lex.Context) (data.LispValue, error) {
    b, ok := ctx.GetByte()
    if !ok {
        return data.Nil, errors.New("eof when parsing atom datatype")
    }
    if !b.IsByteColon() {
        return data.Nil, errors.New("invalid entry point for type atom")
    }
    begin := ctx.Index() + 1
    for {
        ctx.Increment()
        b, ok := ctx.GetByte()
        if !ok {
            return data.Nil, errors.New("premature eof when lexing atom")
        }
        if !(b.IsByteNumber() || b.IsByteLetter() || b.IsByteSpecialSymbol()) {
            return types.NewAtom(ctx.Slice(begin, ctx.Index())), nil
        }
    }
}
