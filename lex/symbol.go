package lex

import (
    "github.com/lucasew/golisp/datatypes"
    "errors"
    // "fmt"
)

func (ctx *Context) ParseSymbol() (datatypes.LispValue, error) {
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
            str := ctx.data[begin:ctx.Index()]
            return datatypes.NewSymbol(string(str)), nil
        }
    }
}
