package lex

import (
    "github.com/lucasew/golisp/datatypes"
    "math/big"
    "errors"
    "fmt"
)

func (ctx *Context) ParseSpecialLiteral() (datatypes.LispValue, error) {
    b, ok := ctx.GetByte()
    if !ok {
        return datatypes.Nil, errors.New("eof when parsing special literal")
    }
    if !b.IsHash() {
        return datatypes.Nil, errors.New("invalid entry point for special literal")
    }
    ctx.Increment()
    b, ok = ctx.GetByte()
    if !ok {
        return datatypes.Nil, errors.New("eof when parsing special literal command")
    }
    if b.IsByteUnderline() { // Like a comment
        ctx.Increment()
        _, err := ctx.ParseString()
        if err != nil {
            return datatypes.Nil, err
        }
        return datatypes.Comment, nil // TODO: Comment object?
    }
    if b.IsByte('b') { // parse binary
        ctx.Increment()
        begin := ctx.Index()
        for {
            b, ok = ctx.GetByte()
            if !ok {
                return datatypes.Nil, errors.New("eof when parsing special literal body")
            }
            if !(b.IsByte('1') || b.IsByte('0')) {
                ret := &big.Int{}
                str := ctx.data[begin:ctx.Index()]
                ret.SetString(string(str), 2) // Parse in base 2
                return datatypes.NewIntFromBigInt(ret), nil
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
                return datatypes.Nil, errors.New("eof when parsing special literal body")
            }
            if !(b >= '0' && b <= '8') {
                ret := &big.Int{}
                str := ctx.data[begin:ctx.Index()]
                ret.SetString(string(str), 8) // Parse in base 8
                return datatypes.NewIntFromBigInt(ret), nil
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
                return datatypes.Nil, errors.New("eof when parsing special literal body")
            }
            if !b.IsHexadecimal() {
                ret := &big.Int{}
                str := ctx.data[begin:ctx.Index()]
                ret.SetString(string(str), 16) // Parse in base 8
                return datatypes.NewIntFromBigInt(ret), nil
            }
            ctx.Increment()
        }
    }
    return datatypes.Nil, fmt.Errorf("i do not understand this special literal expression: '%s'", string(b))
}
