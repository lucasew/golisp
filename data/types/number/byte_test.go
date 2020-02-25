package number

import (
	"github.com/lucasew/golisp/data/types/test"
	"testing"
)

func TestByteType(t *testing.T) {
	v := NewByte(0)
	t.Run("lisp_value", test.NewTestHelper(test.IsValue)(v))
	t.Run("byte", test.NewTestHelper(IsByte)(v))
}
