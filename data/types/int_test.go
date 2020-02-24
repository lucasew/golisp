package types

import (
	"github.com/lucasew/golisp/data/types/number"
	"testing"
)

func TestIntType(t *testing.T) {
	f := number.NewIntFromInt64(0)
	t.Run("lisp_value", ValueTest(f))
	t.Run("number", NumberTest(f))
	t.Run("int", IntTest(f))
}

var IntTest = NewTestHelper(number.IsInt)
