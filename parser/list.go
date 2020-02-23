package parser

import (
	"fmt"
	"github.com/lucasew/golisp/data"
	"github.com/lucasew/golisp/data/types"
	"github.com/lucasew/golisp/lex"
)

func ParseList(ctx *lex.Context, global GlobalStateFunc) (data.LispValue, error) {
	ctx.StateWhitespace()
	b, ok := ctx.GetByte()
	if !ok {
		return types.Nil, fmt.Errorf("%w: list", ErrEOFWhile)
	}
	if !b.IsOpenPar() {
		return types.Nil, fmt.Errorf("%w: list", ErrInvalidEntryPoint)
	}
	ctx.Increment()
	li := types.NewCons()
	for {
		err := ctx.StateWhitespace()
		if err != nil {
			return types.Nil, err
		}
		b, ok = ctx.GetByte()
		if !ok {
			return types.Nil, fmt.Errorf("%w: list", ErrPrematureEOF)
		}
		if b.IsClosePar() {
			ctx.Increment()
			return li, nil
		}
		v, err := global(ctx)
		if err != nil {
			return types.Nil, err
		}
		if !IsComment(v) { // Ignore all comments
			li = append(li.(types.Cons), v)
		}
	}
}
