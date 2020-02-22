package types

import (
    "testing"
)

func TestNilType(t *testing.T) {
    t.Run("lisp_value", ValueTest(Nil))
    t.Run("carcdr", CarCdrTest(Nil))
}
