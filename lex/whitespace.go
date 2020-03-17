package lex

import (
	"fmt"
	"github.com/lucasew/golisp/data"
)

func StateWhitespace(ctx ParseContext) error {
	for {
		select {
		case _ = <-ctx.Done():
			return data.ErrContextCancelled
		default:
			b, ok := ctx.Lex().GetByte()
			if !ok {
				return fmt.Errorf("%w: while looking for whitespaces", ErrEOF)
			}
			if !b.IsBlank() {
				return nil
			}
			ctx.Lex().Increment()
		}
	}
}
