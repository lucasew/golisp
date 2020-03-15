package types

import (
	"github.com/lucasew/golisp/data/entity/test"
	"testing"
)

func TestNilType(t *testing.T) {
	v := Nil
	test.TestValues(v, t, "lisp_carcdr", "lisp_iterator", "lisp_function")
}
