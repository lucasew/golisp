package number

import (
	"github.com/lucasew/golisp/data/types/test"
	"testing"
)

func TestRatType(t *testing.T) {
	v, ok := NewRationalFromString("0")
    if !ok {
        t.Fail()
    }
	t.Run("lisp_value", test.NewTestHelper(test.IsValue)(v))
	t.Run("rational", test.NewTestHelper(IsRational)(v))
}
