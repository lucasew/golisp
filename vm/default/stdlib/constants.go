package stdlib

import (
    "github.com/lucasew/golisp/data/types"
)

func init() {
    register("nil", types.Nil)
    register("t", types.T)
    register("char-bell", "\a")
    register("char-cr", "\r")
    register("char-nl", "\n")
}
