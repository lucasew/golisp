package pdefault

import (
	"fmt"
	"github.com/lucasew/golisp/data"
	"github.com/lucasew/golisp/data/types"
	"github.com/lucasew/golisp/data/types/number"
	"github.com/lucasew/golisp/lex"
	"github.com/lucasew/golisp/parser"
	"math/big"
)

func ParseSpecialLiteral(ctx *lex.Context) (data.LispValue, error) {
	b, ok := ctx.GetByte()
	if !ok {
		return types.Nil, fmt.Errorf("%w: special literal", parser.ErrEOFWhile)
	}
	if !b.IsHash() {
		return types.Nil, fmt.Errorf("%w: special literal", parser.ErrInvalidEntryPoint)
	}
	ctx.Increment()
	b, ok = ctx.GetByte()
	if !ok {
		return types.Nil, fmt.Errorf("%w: special literal command", parser.ErrPrematureEOF)
	}
	if b.IsByteUnderline() { // Like a comment
		ctx.Increment()
		_, err := parser.ParseString(ctx)
		if err != nil {
			return types.Nil, err
		}
		return parser.Comment, nil // TODO: Comment object?
	}
	if b.IsByte('b') { // parse binary
		ctx.Increment()
		begin := ctx.Index()
		for {
			b, ok = ctx.GetByte()
			if !ok {
				return types.Nil, fmt.Errorf("%w: special literal body", parser.ErrPrematureEOF)
			}
			if !(b.IsByte('1') || b.IsByte('0')) {
				ret := &big.Int{}
				s := ctx.Slice(begin, ctx.Index())
				ret.SetString(s, 2) // Parse in base 2
				return number.NewIntFromBigInt(ret), nil
			}
			ctx.Increment()
		}
	}
	if b.IsByte('o') { // parse octal
		ctx.Increment()
		begin := ctx.Index()
		for {
			b, ok = ctx.GetByte()
			if !ok {
				return types.Nil, fmt.Errorf("%w: special literal body", parser.ErrPrematureEOF)
			}
			if !(b >= '0' && b <= '8') {
				ret := &big.Int{}
				s := ctx.Slice(begin, ctx.Index())
				ret.SetString(s, 8) // Parse in base 8
				return number.NewIntFromBigInt(ret), nil
			}
			ctx.Increment()
		}
	}
	if b.IsByte('x') { // parse octal
		ctx.Increment()
		begin := ctx.Index()
		for {
			b, ok = ctx.GetByte()
			if !ok {
				return types.Nil, fmt.Errorf("%w: special literal body", parser.ErrPrematureEOF)
			}
			if !b.IsHexadecimal() {
				ret := &big.Int{}
				s := ctx.Slice(begin, ctx.Index())
				ret.SetString(s, 16) // Parse in base 8
				return number.NewIntFromBigInt(ret), nil
			}
			ctx.Increment()
		}
	}
	if b.IsByte('c') { // parse byte/char
		ctx.Increment()
		n, err := parser.ParseNumber(ctx)
		if err != nil {
			return types.Nil, err
		}
		num, ok := n.(number.LispInt)
		if !ok {
			return types.Nil, fmt.Errorf("invalid syntax for byte parsing")
		}
		inum, _ := num.Int64()
		return number.NewByte(byte(inum)), nil
	}
	return types.Nil, fmt.Errorf("%w: i do not understand this special literal expression: '%s'", parser.ErrInvalidChar, string(b))
}
