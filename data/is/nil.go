package is

import (
    "github.com/lucasew/golisp/data"
)

func Nil(v data.LispValue) bool {
    return v.IsNil()
}
