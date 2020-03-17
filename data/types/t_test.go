package types

import (
	"github.com/lucasew/golisp/data/entity/test"
	"testing"
)

func TestTType(t *testing.T) {
	v := T
	test.TestValues(v, t, "lisp_carcdr", "lisp_function")
}
