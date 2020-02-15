package parser

import (
    "github.com/lucasew/golisp/datatypes"
    "github.com/lucasew/golisp/lex"
)

type ParserFunc func(string) (datatypes.LispValue, error)

type GlobalStateFunc func(*lex.Context) (datatypes.LispValue, error)
