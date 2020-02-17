package parser

import (
    "github.com/lucasew/golisp/data"
)

type _comment uint8

var Comment data.LispValue = _comment(0)

func (_comment) Car() data.LispValue {
    return nil
}

func (_comment) Cdr() data.LispValue {
    return nil
}

func (_comment) IsNil() bool {
    return true
}

func (_comment) Repr() string {
    return ""
}

func IsComment(lv data.LispValue) bool {
    _, ok := lv.(_comment)
    return ok
}
