package types

import (
	"testing"
)

func TestTType(t *testing.T) {
	t.Run("lisp_value", ValueTest(T))
	t.Run("carcdr", CarCdrTest(T))
}
