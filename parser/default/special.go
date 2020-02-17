package pdefault

import (
    "github.com/lucasew/golisp/data"
    "github.com/lucasew/golisp/data/types"
    "github.com/lucasew/golisp/lex"
    "github.com/lucasew/golisp/parser"
    "math/big"
    "errors"
    "fmt"
)

func ParseSpecialLiteral(ctx *lex.Context) (data.LispValue, error) {
    b, ok := ctx.GetByte()
    if !ok {
        return data.Nil, errors.New("eof when parsing special literal")
    }
    if !b.IsHash() {
        return data.Nil, errors.New("invalid entry point for special literal")
    }
    ctx.Increment()
    b, ok = ctx.GetByte()
    if !ok {
        return data.Nil, errors.New("eof when parsing special literal command")
    }
    if b.IsByteUnderline() { // Like a comment
        ctx.Increment()
        _, err := parser.ParseString(ctx)
        if err != nil {
            return data.Nil, err
        }
        return parser.Comment, nil // TODO: Comment object?
    }
    if b.IsByte('b') { // parse binary
        ctx.Increment()
        begin := ctx.Index()
        for {
            b, ok = ctx.GetByte()
            if !ok {
                return data.Nil, errors.New("eof when parsing special literal body")
            }
            if !(b.IsByte('1') || b.IsByte('0')) {
                ret := &big.Int{}
                s := ctx.Slice(begin, ctx.Index())
                ret.SetString(s, 2) // Parse in base 2
                return types.NewIntFromBigInt(ret), nil
            }
            ctx.Increment()
        }
    }
    if b.IsByte('o') { // parse octal
        ctx.Increment()
        begin := ctx.Index()
        for {
            b, ok = ctx.GetByte()
            if !ok {
                return data.Nil, errors.New("eof when parsing special literal body")
            }
            if !(b >= '0' && b <= '8') {
                ret := &big.Int{}
                s := ctx.Slice(begin, ctx.Index())
                ret.SetString(s, 8) // Parse in base 8
                return types.NewIntFromBigInt(ret), nil
            }
            ctx.Increment()
        }
    }
    if b.IsByte('x') { // parse octal
        ctx.Increment()
        begin := ctx.Index()
        for {
            b, ok = ctx.GetByte()
            if !ok {
                return data.Nil, errors.New("eof when parsing special literal body")
            }
            if !b.IsHexadecimal() {
                ret := &big.Int{}
                s := ctx.Slice(begin, ctx.Index())
                ret.SetString(s, 16) // Parse in base 8
                return types.NewIntFromBigInt(ret), nil
            }
            ctx.Increment()
        }
    }
    return data.Nil, fmt.Errorf("i do not understand this special literal expression: '%s'", string(b))
}
