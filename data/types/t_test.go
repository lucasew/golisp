package types

import (
	"testing"
    "github.com/lucasew/golisp/data/types/test"
)

func TestTType(t *testing.T) {
    v := T
	t.Run("lisp_value", test.NewTestHelper(test.IsValue)(v))
	t.Run("carcdr", test.NewTestHelper(test.IsCarCdr)(v))
}
