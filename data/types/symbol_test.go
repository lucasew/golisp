package types

import (
	"testing"
)

func TestSymbolType(t *testing.T) {
	f := NewSymbol("")
	t.Run("lisp_value", ValueTest(f))
	t.Run("string", StringTest(f))
	t.Run("symbol", SymbolTest(f))
}

var SymbolTest = NewTestHelper(IsSymbol)
