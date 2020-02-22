package is

import (
    "github.com/lucasew/golisp/data"
)

func Value(v data.LispValue) bool {
    _, ok := v.(data.LispValue)
    return ok
}
