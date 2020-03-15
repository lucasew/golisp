package types

import (
	"github.com/lucasew/golisp/data/entity/test"
	"testing"
)

func TestFunctionType(t *testing.T) {
	v := NewFunction(nil)
	test.TestValues(v, t, "lisp_function")
}
