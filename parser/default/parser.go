package pdefault

import (
	"context"
	"github.com/lucasew/golisp/data"
	"github.com/lucasew/golisp/data/types"
	"github.com/lucasew/golisp/lex"
	"github.com/lucasew/golisp/parser"
)

func Parse(ctx context.Context, s string) (data.LispValue, error) {
	pctx := lex.NewParseContextFromContext(ctx, lex.NewContext([]byte(s)))
	list := types.NewCons().(types.Cons)
	err := error(nil)
	for {
		ret, err := GlobalState(pctx)
		if err != nil {
			return types.Nil, err
		}
		if !parser.IsComment(ret) { // Ignore all comments
			list = append(list, ret)
		}
		err = lex.StateWhitespace(pctx)
		if err != nil {
			break
		}
	}
	return list, err
}
