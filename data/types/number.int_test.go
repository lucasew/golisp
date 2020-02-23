package types

import (
	"testing"
)

func TestIntType(t *testing.T) {
	f := NewIntFromInt64(0)
	t.Run("lisp_value", ValueTest(f))
	t.Run("number", NumberTest(f))
	t.Run("int", IntTest(f))
}

var IntTest = NewTestHelper(IsInt)
