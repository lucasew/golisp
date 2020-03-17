package parser

import (
	"fmt"
	"github.com/lucasew/golisp/data"
	"github.com/lucasew/golisp/data/types"
	"github.com/lucasew/golisp/lex"
)

func ParseSymbol(ctx lex.ParseContext) (data.LispValue, error) {
	b, ok := ctx.Lex().GetByte()
	if !ok {
		return types.Nil, fmt.Errorf("%w: symbol", ErrEOFWhile)
	}
	if !(b.IsByteLetter() || b.IsByteSpecialSymbol()) {
		return types.Nil, fmt.Errorf("%w: symbol", ErrInvalidEntryPoint)
	}
	begin := ctx.Lex().Index()
	for {
		ctx.Lex().Increment()
		b, ok = ctx.Lex().GetByte()
		if !ok {
			return types.Nil, fmt.Errorf("%w: symbol", ErrPrematureEOF)
		}
		if !(b.IsByteLetter() || b.IsByteSpecialSymbol()) {
			s := ctx.Lex().Slice(begin, ctx.Lex().Index())
			return types.NewSymbol(s), nil
		}
	}
}
