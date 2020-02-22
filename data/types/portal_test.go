
package types

import (
    "testing"
)

func TestPortalType(t *testing.T) {
    f := NewPortal(0)
    t.Run("lisp_value", ValueTest(f))
    t.Run("portal", PortalTest(f))
}

var PortalTest = NewTestHelper(IsPortal)
