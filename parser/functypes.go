package parser

import (
	"context"
	"github.com/lucasew/golisp/data"
	"github.com/lucasew/golisp/lex"
)

type ParserFunc func(context.Context, string) (data.LispValue, error)

type GlobalStateFunc func(lex.ParseContext) (data.LispValue, error)
