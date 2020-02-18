package stdlib

import (
    "github.com/lucasew/golisp/data"
)

func init() {
    register("nil", data.Nil)
    register("t", data.T)
    register("char-bell", "\a")
    register("char-cr", "\r")
    register("char-nl", "\n")
}
