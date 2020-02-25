package number

import (
	"github.com/lucasew/golisp/data/types/test"
	"testing"
)

func TestFloatType(t *testing.T) {
	v := NewFloatFromFloat64(0)
	t.Run("lisp_value", test.NewTestHelper(test.IsValue)(v))
	t.Run("number", test.NewTestHelper(test.IsNumber)(v))
	t.Run("float", test.NewTestHelper(IsFloat)(v))
}
