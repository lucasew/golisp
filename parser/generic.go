package parser

import (
    "github.com/lucasew/golisp/data"
    "github.com/lucasew/golisp/lex"
)

type ParserFunc func(string) (data.LispValue, error)

type GlobalStateFunc func(*lex.Context) (data.LispValue, error)
