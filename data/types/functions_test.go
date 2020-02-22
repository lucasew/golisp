package types

import (
    "testing"
)

func TestFunctionType(t *testing.T) {
    fn := NewFunction(nil)
    t.Run("lisp_value", ValueTest(fn))
    t.Run("function", FunctionTest(fn))
}
