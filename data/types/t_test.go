package types

import (
	"github.com/lucasew/golisp/data/types/test"
	"testing"
)

func TestTType(t *testing.T) {
	v := T
	t.Run("lisp_value", test.NewTestHelper(test.IsValue)(v))
	t.Run("carcdr", test.NewTestHelper(test.IsCarCdr)(v))
	t.Run("function", test.NewTestHelper(test.IsFunction)(v))
}
