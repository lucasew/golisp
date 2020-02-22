package is

import (
    "github.com/lucasew/golisp/data"
)

func CarCdr(v data.LispValue) bool {
    _, ok := v.(data.LispCarCdr)
    return ok
}
