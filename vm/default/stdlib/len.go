package stdlib

import (
    "github.com/lucasew/golisp/data"
    "github.com/lucasew/golisp/data/convert"
)

func init() {
    register("len", Len)
}

func Len(v data.LispCons) (data.LispValue, error) {
    return convert.NewLispValue(_len(v.Car()))
}

func _len(v data.LispValue) int {
    l, ok := v.(data.LispLen)
    if ok {
        return l.Len()
    } else {
        if v.IsNil() {
            return 0
        } else {
            return 1
        }
    }
}
