package types

import (
    "testing"
)

func TestByteType(t *testing.T) {
    f := NewByte(0)
    t.Run("lisp_value", ValueTest(f))
    t.Run("byte", ByteTest(f))
}
