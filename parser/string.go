package parser

import (
	"fmt"
	"github.com/lucasew/golisp/data"
	"github.com/lucasew/golisp/data/types"
	"github.com/lucasew/golisp/lex"
)

func ParseString(ctx *lex.Context) (data.LispValue, error) {
	b, ok := ctx.GetByte()
	if !ok {
		return types.Nil, fmt.Errorf("%w: string", ErrEOFWhile)
	}
	if !b.IsStringMark() {
		return types.Nil, fmt.Errorf("%w: string", ErrInvalidEntryPoint)
	}
	begin := ctx.Index() + 1
	for {
		ctx.Increment()
		b, ok := ctx.GetByte()
		if !ok {
			return types.Nil, fmt.Errorf("%w: string", ErrPrematureEOF)
		}
		if b.IsStringMark() {
			s := ctx.Slice(begin, ctx.Index())
			ctx.Increment()
			return types.NewString(s), nil
		}
	}
}
