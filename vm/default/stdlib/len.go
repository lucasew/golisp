package stdlib

import (
    "github.com/lucasew/golisp/data"
    "github.com/lucasew/golisp/data/convert"
)

func init() {
    register("len", Len)
}

func Len(v data.LispValue) (data.LispValue, error) {
    return convert.NewLispValue(_len(v.Car()))
}

func _len(v data.LispValue) int {
    if v.IsNil() {
        return 0
    } else {
        return 1 + _len(v.Cdr())
    }
}
