package types

import (
    "testing"
)

func TestRatType(t *testing.T) {
    f, _ := NewRationalFromString("0")
    t.Run("lisp_value", ValueTest(f))
    t.Run("byte", RationalTest(f))
}
