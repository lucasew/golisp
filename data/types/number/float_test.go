package number

import (
	"github.com/lucasew/golisp/data/entity/test"
	"testing"
)

func TestFloatType(t *testing.T) {
	v := NewFloatFromFloat64(0)
	test.TestValues(v, t, "lisp_number", "float")
}
