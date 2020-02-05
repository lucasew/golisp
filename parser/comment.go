package parser

import (
    "github.com/lucasew/golisp/datatypes"
)

type _comment uint8

var Comment datatypes.LispValue = _comment(0)

func (_comment) Car() datatypes.LispValue {
    return nil
}

func (_comment) Cdr() datatypes.LispValue {
    return nil
}

func (_comment) IsNil() bool {
    return true
}

func (_comment) Repr() string {
    return ""
}

func IsComment(lv datatypes.LispValue) bool {
    _, ok := lv.(_comment)
    return ok
}
