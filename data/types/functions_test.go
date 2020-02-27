package types

import (
	"github.com/lucasew/golisp/data/types/test"
	"testing"
)

func TestFunctionType(t *testing.T) {
	v := NewFunction(nil)
	t.Run("lisp_value", test.NewTestHelper(test.IsValue)(v))
	t.Run("function", test.NewTestHelper(test.IsFunction)(v))
}
