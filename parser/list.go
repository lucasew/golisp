package parser

import (
    "github.com/lucasew/golisp/data"
    "github.com/lucasew/golisp/data/types"
    "github.com/lucasew/golisp/lex"
    "fmt"
)

func ParseList(ctx *lex.Context, global GlobalStateFunc) (data.LispValue, error) {
    ctx.StateWhitespace()
    b, ok := ctx.GetByte()
    if !ok {
        return data.Nil, fmt.Errorf("%w: list", ErrEOFWhile)
    }
    if !b.IsOpenPar() {
        return data.Nil, fmt.Errorf("%w: list", ErrInvalidEntryPoint)
    }
    ctx.Increment()
    li := types.NewCons()
    for {
        err := ctx.StateWhitespace()
        if err != nil {
            return data.Nil, err
        }
        b, ok = ctx.GetByte()
        if !ok {
            return data.Nil, fmt.Errorf("%w: list", ErrPrematureEOF)
        }
        if b.IsClosePar() {
            ctx.Increment()
            return li, nil
        }
        v, err := global(ctx)
        if err != nil {
            return data.Nil, err
        }
        if !IsComment(v) { // Ignore all comments
            li = append(li.(types.Cons), v)
        }
    }
}
