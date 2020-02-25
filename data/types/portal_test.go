package types

import (
	"testing"
    "github.com/lucasew/golisp/data/types/test"
)

func TestPortalType(t *testing.T) {
	v := NewPortal(0)
	t.Run("lisp_value", test.NewTestHelper(test.IsValue)(v))
	t.Run("portal", test.NewTestHelper(IsPortal)(v))
}
