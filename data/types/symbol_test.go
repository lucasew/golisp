package types

import (
	"github.com/lucasew/golisp/data/entity/test"
	_ "github.com/lucasew/golisp/data/types/test"
	"testing"
)

func TestSymbolType(t *testing.T) {
	v := NewSymbol("")
	test.TestValues(v, t, "lisp_string", "lisp_carcdr", "symbol")
}
