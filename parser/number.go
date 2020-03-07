package parser

import (
	"fmt"
	"github.com/lucasew/golisp/data"
	"github.com/lucasew/golisp/data/types"
	"github.com/lucasew/golisp/data/types/number"
	"github.com/lucasew/golisp/lex"
	"strings"
)

func ParseNumber(ctx *lex.Context) (data.LispValue, error) {
	b, ok := ctx.GetByte()
	if !ok {
		return types.Nil, fmt.Errorf("%w: number", ErrEOFWhile)
	}
	if !(b.IsByteNumber() || b.IsByte('-') || b.IsByte('.')) {
		return types.Nil, fmt.Errorf("%w: number", ErrInvalidEntryPoint)
	}
	e := false     // For scientific notation like 10^2 that is like 1E2
	dot := false   // For the dot of floats
	slash := false // For the / in rational numbers
	begin := ctx.Index()
	for {
		ctx.Increment()
		b, ok := ctx.GetByte()
		if !ok {
			// return types.Nil, errors.New("eof when parsing number body")
		}
		if b.IsByteE() {
			if e {
				return types.Nil, fmt.Errorf("%w: 'e'", ErrInvalidMarkerRepetition)
			}
			e = true
			continue
		}
		if b.IsDot() {
			if dot {
				return types.Nil, fmt.Errorf("%w: '.'", ErrInvalidMarkerRepetition)
			}
			dot = true
			continue
		}
		if b.IsSlash() {
			if slash {
				return types.Nil, fmt.Errorf("%w: '/'", ErrInvalidMarkerRepetition)
			}
			slash = true
			continue
		}
		if b.IsByte('-') {
			continue
		}
		if b.IsByteNumber() {
			continue
		}
		if b.IsByteUnderline() { // TODO: Ignore it when using the function to parse
			continue
		}
		s := ctx.Slice(begin, ctx.Index())
		s = strings.ReplaceAll(s, "_", "")
		reti, ok := number.NewIntFromString(s)
		if ok {
			return reti, nil
		}
		retf, ok := number.NewFloatFromString(s)
		if ok {
			return retf, nil
		}

		retr, ok := number.NewRationalFromString(s)
		if ok {
			return retr, nil
		}
		return types.Nil, fmt.Errorf("%w: %s", ErrCantParseAsNumber, s)
	}
}
