package parser

import (
    "github.com/lucasew/golisp/datatypes"
    "github.com/lucasew/golisp/lex"
)

type GlobalStateFunc func(*lex.Context) (datatypes.LispValue, error)
