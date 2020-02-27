package types

import (
	"github.com/lucasew/golisp/data/types/test"
	"testing"
)

func TestNilType(t *testing.T) {
	v := Nil
	t.Run("lisp_value", test.NewTestHelper(test.IsValue)(v))
	t.Run("carcdr", test.NewTestHelper(test.IsCarCdr)(v))
	t.Run("iterator", test.NewTestHelper(test.IsIterator)(v))
	t.Run("function", test.NewTestHelper(test.IsFunction)(v))
}
