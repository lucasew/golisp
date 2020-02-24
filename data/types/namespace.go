package types

import (
    "github.com/lucasew/golisp/data"
)

func IsNamespace(v data.LispValue) bool {
    _, ok := v.(data.LispNamespace)
    return ok
}
