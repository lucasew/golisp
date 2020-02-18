package lex

import (
    "fmt"
    // "github.com/davecgh/go-spew/spew"
)

func (ctx *Context) StateWhitespace() error {
    for {
        b, ok := ctx.GetByte()
        if !ok {
            return fmt.Errorf("%w: while looking for whitespaces", ErrEOF)
        }
        if !b.IsBlank() {
            return nil
        }
        ctx.Increment()
    }
}
