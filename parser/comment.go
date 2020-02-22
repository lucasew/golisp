package parser

import (
    "github.com/lucasew/golisp/data"
)

type _comment uint8

var Comment data.LispValue = _comment(0)

func (_comment) IsNil() bool {
    return true
}

func (_comment) Repr() string {
    return ""
}

func (_comment) LispTypeName() string {
    return "comment"
}

func IsComment(lv data.LispValue) bool {
    _, ok := lv.(_comment)
    return ok
}

