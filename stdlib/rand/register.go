package rand

import (
    "github.com/lucasew/golisp/stdlib/loader"
)

var ELEMENTS = loader.NewRepository()

func register(k string, v interface{}) {
    ELEMENTS.Register("rand", k, func() interface{} {return v})
}
