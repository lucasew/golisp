package types

import (
	"testing"
	"github.com/lucasew/golisp/data/types/test"
)

func TestNilType(t *testing.T) {
    v := Nil
	t.Run("lisp_value", test.NewTestHelper(test.IsValue)(v))
	t.Run("carcdr", test.NewTestHelper(test.IsCarCdr)(v))
    t.Run("iterator", test.NewTestHelper(test.IsIterator)(v))
}
