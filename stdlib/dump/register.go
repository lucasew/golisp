package libdump

import (
	"github.com/lucasew/golisp/data"
	"github.com/lucasew/golisp/data/convert"
)

var ELEMENTS = map[string]data.LispValue{}

func register(k string, v interface{}) {
	var err error
	ELEMENTS[k], err = convert.NewLispValue(v)
	if err != nil {
		panic(err)
	}
}
