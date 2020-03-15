package rand

import (
	"github.com/lucasew/golisp/stdlib"
)

var ELEMENTS = stdlib.NewRepository()

func register(k string, v interface{}) {
	ELEMENTS.Register("rand", k, func() interface{} { return v })
}
