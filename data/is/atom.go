package is

import (
    "github.com/lucasew/golisp/data"
)

func Atom(v data.LispValue) bool {
    _, ok := v.(data.LispAtom)
    return ok
}
