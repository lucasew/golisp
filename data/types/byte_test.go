package types

import (
	"testing"
    "github.com/lucasew/golisp/data/types/number"
)

func TestByteType(t *testing.T) {
	f := number.NewByte(0)
	t.Run("lisp_value", ValueTest(f))
	t.Run("byte", ByteTest(f))
}

var ByteTest = NewTestHelper(number.IsByte)
