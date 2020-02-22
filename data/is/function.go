package is

import (
    "github.com/lucasew/golisp/data"
)

func Function(v data.LispValue) bool {
    _, ok := v.(data.LispFunction)
    return ok
}
