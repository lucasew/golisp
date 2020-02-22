package is

import (
    "github.com/lucasew/golisp/data"
)

func Cons(v data.LispValue) bool {
    _, ok := v.(data.LispCons)
    return ok
}
