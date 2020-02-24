package types

import (
	"github.com/lucasew/golisp/data/types/number"
	"testing"
)

func TestFloatType(t *testing.T) {
	f := number.NewFloatFromFloat64(0)
	t.Run("lisp_value", ValueTest(f))
	t.Run("number", NumberTest(f))
	t.Run("float", FloatTest(f))
}

var FloatTest = NewTestHelper(number.IsFloat)
