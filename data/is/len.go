package is

import (
    "github.com/lucasew/golisp/data"
)

func Len(v data.LispValue) bool {
    _, ok := v.(data.LispLen)
    return ok
}
