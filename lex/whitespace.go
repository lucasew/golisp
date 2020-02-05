package lex

import (
    "errors"
    // "github.com/davecgh/go-spew/spew"
)

func (ctx *Context) StateWhitespace() error {
    for {
        b, ok := ctx.GetByte()
        if !ok {
            return errors.New("eof while looking for whitespaces")
        }
        if !b.IsBlank() {
            return nil
        }
        ctx.Increment()
    }
}
