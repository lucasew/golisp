package types

import (
	"testing"
    "github.com/lucasew/golisp/data/types/number"
)

func TestRatType(t *testing.T) {
	f, _ := number.NewRationalFromString("0")
	t.Run("lisp_value", ValueTest(f))
	t.Run("byte", RationalTest(f))
}

var RationalTest = NewTestHelper(number.IsRational)
