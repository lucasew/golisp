package parser

import (
	"fmt"
	"github.com/lucasew/golisp/data"
	"github.com/lucasew/golisp/data/types"
	"github.com/lucasew/golisp/lex"
)

func ParseAtom(ctx *lex.Context) (data.LispValue, error) {
	b, ok := ctx.GetByte()
	if !ok {
		return types.Nil, fmt.Errorf("%w: atom datatype", ErrEOFWhile)
	}
	if !b.IsByteColon() {
		return types.Nil, fmt.Errorf("%w: atom datatype", ErrInvalidEntryPoint)
	}
	begin := ctx.Index() + 1
	for {
		ctx.Increment()
		b, ok := ctx.GetByte()
		if !ok {
			return types.Nil, fmt.Errorf("%w: atom", ErrPrematureEOF)
		}
		if !(b.IsByteNumber() || b.IsByteLetter() || b.IsByteSpecialSymbol()) {
			return types.NewAtom(ctx.Slice(begin, ctx.Index())), nil
		}
	}
}
