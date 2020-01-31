package lex

import (
    "github.com/lucasew/golisp/datatypes"
    "errors"
)

func (ctx *Context) ParseString() (datatypes.LispValue, error) {
    b, ok := ctx.GetByte()
    if !ok {
        return datatypes.Nil, errors.New("eof when parsing string")
    }
    if !b.IsStringMark() {
        return datatypes.Nil, errors.New("invalid entry point for string")
    }
    begin := ctx.Index() + 1
    for {
        ctx.Increment()
        b, ok := ctx.GetByte()
        if !ok {
            return datatypes.Nil, errors.New("eof when parsing string body")
        }
        if b.IsStringMark() {
            block := ctx.data[begin:ctx.Index()]
            ctx.Increment()
            return datatypes.NewConventionalString(string(block)), nil
        }
    }
}
