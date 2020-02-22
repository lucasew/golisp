package is

import (
    "github.com/lucasew/golisp/data"
)

func NativeFunction(v data.LispValue) bool {
    fn, ok := v.(data.LispFunction)
    if !ok {
        return false
    }
    return fn.IsFunctionNative()
}
