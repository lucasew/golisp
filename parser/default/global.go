package pdefault

import (
	"fmt"
	"github.com/lucasew/golisp/data"
	"github.com/lucasew/golisp/data/types"
	"github.com/lucasew/golisp/lex"
	"github.com/lucasew/golisp/parser"
)

func GlobalState(ctx lex.ParseContext) (data.LispValue, error) {
	err := lex.StateWhitespace(ctx)
	if err != nil {
		return types.Nil, err
	}
	select {
	case _ = <-ctx.Done():
		return types.Nil, data.ErrContextCancelled
	default:
		b, ok := ctx.Lex().GetByte()
		if !ok {
			return types.Nil, fmt.Errorf("%w: global state", parser.ErrEOFWhile)
		}
		if b.IsByte('-') {
			ctx.Lex().Increment()
			b, ok = ctx.Lex().GetByte()
			if !ok {
				return types.Nil, fmt.Errorf("%w: comment marker", parser.ErrEOFWhile)
			}
			if b.IsByteNumber() {
				ctx.Lex().Decrement()
				return parser.ParseNumber(ctx)
			}
			if b.IsBlank() {
				return types.NewSymbol("-"), nil
			}
			if b.IsByte('-') {
				ctx.Lex().Increment()
				for {
					b, ok := ctx.Lex().GetByte()
					if !ok {
						return types.Nil, fmt.Errorf("%w: comment", parser.ErrPrematureEOF)
					}
					if b.IsByte('\n') {
						return parser.Comment, nil
					}
					ctx.Lex().Increment()
				}
			} else {
				sym, err := parser.ParseSymbol(ctx)
				if err != nil {
					return types.Nil, err
				}
				s := fmt.Sprintf("-%s", sym)
				return types.NewSymbol(s), nil
			}
		}
		if b.IsByte('\'') {
			ctx.Lex().Increment()
			g, err := GlobalState(ctx)
			return types.NewCons(types.NewSymbol("quote"), g), err
		}
		if b.IsByteColon() {
			return parser.ParseAtom(ctx)
		}
		if b.IsOpenPar() {
			return parser.ParseList(ctx, GlobalState)
		}
		if b.IsStringMark() {
			return parser.ParseString(ctx)
		}
		if b.IsByteNumber() || b.IsByte('.') {
			return parser.ParseNumber(ctx)
		}
		if b.IsHash() {
			return ParseSpecialLiteral(ctx)
		}
		if b.IsByteLetter() || b.IsByteSpecialSymbol() {
			return parser.ParseSymbol(ctx)
		}
		// if b.IsClosePar() { // TODO: Test more
		//     // panic("invalid ) token")
		//     return types.Nil, errors.New("invalid ')' token")
		//     // return types.Nil, nil
		// }
		return types.Nil, fmt.Errorf("%w: '%s'", parser.ErrInvalidChar, string(ctx.Lex().MustGetByte()))
	}
}
