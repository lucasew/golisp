package number

import (
	"github.com/lucasew/golisp/data/types/test"
	"testing"
)

func TestIntType(t *testing.T) {
	v := NewIntFromInt64(0)
	t.Run("lisp_value", test.NewTestHelper(test.IsValue)(v))
	t.Run("number", test.NewTestHelper(test.IsNumber)(v))
	t.Run("int", test.NewTestHelper(IsInt)(v))
}
