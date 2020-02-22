package types

import (
    "testing"
)

func TestFloatType(t *testing.T) {
    f := NewFloatFromFloat64(0)
    t.Run("lisp_value", ValueTest(f))
    t.Run("number", NumberTest(f))
    t.Run("float", FloatTest(f))
}
