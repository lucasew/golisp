package types

import (
	"testing"
    "github.com/lucasew/golisp/data/types/test"
)

func TestAtomType(t *testing.T) {
    v := NewAtom("")
	t.Run("lisp_value", test.NewTestHelper(test.IsValue)(v))
	t.Run("atom", test.NewTestHelper(test.IsAtom)(v))
}
