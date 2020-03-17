package types

import (
	"github.com/lucasew/golisp/data/entity/test"
	"testing"
)

func TestPortalType(t *testing.T) {
	v := NewPortal(0)
	test.TestValues(v, t, "lisp_portal", "portal")
}
