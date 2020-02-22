package is

import (
    "github.com/lucasew/golisp/data"
)

func Map(v data.LispValue) bool {
    _, ok := v.(data.LispMap)
    return ok
}
