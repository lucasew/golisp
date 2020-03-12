package libdump

import (
	"github.com/lucasew/golisp/stdlib/loader"
)

var ELEMENTS = loader.NewRepository()

func register(k string, v interface{}) {
	ELEMENTS.Register("dump", k, func() interface{} { return v })
}
