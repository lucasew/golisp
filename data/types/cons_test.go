package types

import (
	"testing"
    "github.com/lucasew/golisp/data/types/test"
)

func TestConsType(t *testing.T) {
    v := NewCons()
	t.Run("lisp_value", test.NewTestHelper(test.IsValue)(v))
	t.Run("is_cons", test.NewTestHelper(test.IsCons)(v))
	t.Run("carcdr", test.NewTestHelper(test.IsCarCdr)(v))
}

