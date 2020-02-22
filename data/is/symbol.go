package is

import (
    "github.com/lucasew/golisp/data"
    "github.com/lucasew/golisp/data/types"
)

func Symbol(v data.LispValue) bool {
    _, ok := v.(types.Symbol)
    return ok
}
