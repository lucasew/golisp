package types

import (
	"github.com/lucasew/golisp/data/types/test"
	"testing"
)

func TestSymbolType(t *testing.T) {
	v := NewSymbol("")
	t.Run("lisp_value", test.NewTestHelper(test.IsValue)(v))
	t.Run("string", test.NewTestHelper(test.IsString)(v))
	t.Run("carcdr", test.NewTestHelper(test.IsCarCdr)(v))
	t.Run("symbol", test.NewTestHelper(IsSymbol)(v))
}
