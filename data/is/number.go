package is

import (
    "github.com/lucasew/golisp/data"
)

func Number(v data.LispValue) bool {
    _, ok := v.(data.LispNumber)
    return ok
}
