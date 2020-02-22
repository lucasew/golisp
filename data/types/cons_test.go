package types

import (
    "testing"
)

func TestConsType(t *testing.T) {
    t.Run("lisp_value", ValueTest(NewCons()))
    t.Run("is_cons", ConsTest(NewCons()))
    t.Run("carcdr", CarCdrTest(NewCons()))
}
