package lex

import (
    "github.com/lucasew/golisp/datatypes"
    "errors"
)

func (ctx *Context) ParseAtom() (datatypes.LispValue, error) {
    b, ok := ctx.GetByte()
    if !ok {
        return datatypes.Nil, errors.New("eof when parsing atom datatype")
    }
    if !b.IsByteColon() {
        return datatypes.Nil, errors.New("invalid entry point for type atom")
    }
    begin := ctx.Index() + 1
    for {
        ctx.Increment()
        b, ok := ctx.GetByte()
        if !ok {
            return datatypes.Nil, errors.New("premature eof when lexing atom")
        }
        if !(b.IsByteNumber() || b.IsByteLetter() || b.IsByteSpecialSymbol()) {
            return datatypes.NewAtom(string(ctx.data[begin:ctx.Index()])), nil
        }
    }
}
