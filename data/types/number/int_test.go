package number

import (
	"github.com/lucasew/golisp/data/entity/test"
	"testing"
)

func TestIntType(t *testing.T) {
	v := NewIntFromInt64(0)
	test.TestValues(v, t, "lisp_number", "int")
}
