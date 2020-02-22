package is

import (
    "github.com/lucasew/golisp/data"
)

func String(v data.LispValue) bool {
    _, ok := v.(data.LispString)
    return ok
}
