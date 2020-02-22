package is

import (
    "github.com/lucasew/golisp/data"
)

func Portal(v data.LispValue) bool {
    _, ok := v.(data.LispPortal)
    return ok
}
