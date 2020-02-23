package parser

import (
	"fmt"
	"github.com/lucasew/golisp/data"
	"github.com/lucasew/golisp/data/types"
	"github.com/lucasew/golisp/lex"
)

func ParseSymbol(ctx *lex.Context) (data.LispValue, error) {
	b, ok := ctx.GetByte()
	if !ok {
		return types.Nil, fmt.Errorf("%w: symbol", ErrEOFWhile)
	}
	if !(b.IsByteLetter() || b.IsByteSpecialSymbol()) {
		return types.Nil, fmt.Errorf("%w: symbol", ErrInvalidEntryPoint)
	}
	begin := ctx.Index()
	for {
		ctx.Increment()
		b, ok = ctx.GetByte()
		if !ok {
			return types.Nil, fmt.Errorf("%w: symbol", ErrPrematureEOF)
		}
		if !(b.IsByteLetter() || b.IsByteSpecialSymbol()) {
			s := ctx.Slice(begin, ctx.Index())
			return types.NewSymbol(s), nil
		}
	}
}
