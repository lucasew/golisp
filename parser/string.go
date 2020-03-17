package parser

import (
	"fmt"
	"github.com/lucasew/golisp/data"
	"github.com/lucasew/golisp/data/types"
	"github.com/lucasew/golisp/lex"
)

func ParseString(ctx lex.ParseContext) (data.LispValue, error) {
	b, ok := ctx.Lex().GetByte()
	if !ok {
		return types.Nil, fmt.Errorf("%w: string", ErrEOFWhile)
	}
	if !b.IsStringMark() {
		return types.Nil, fmt.Errorf("%w: string", ErrInvalidEntryPoint)
	}
	begin := ctx.Lex().Index() + 1
	for {
		ctx.Lex().Increment()
		b, ok := ctx.Lex().GetByte()
		if !ok {
			return types.Nil, fmt.Errorf("%w: string", ErrPrematureEOF)
		}
		if b.IsStringMark() {
			s := ctx.Lex().Slice(begin, ctx.Lex().Index())
			ctx.Lex().Increment()
			return types.NewString(s), nil
		}
	}
}
