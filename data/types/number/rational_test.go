package number

import (
	"github.com/lucasew/golisp/data/entity/test"
	"testing"
)

func TestRatType(t *testing.T) {
	v, ok := NewRationalFromString("0")
	if !ok {
		t.Fail()
	}
	test.TestValues(v, t, "rational", "lisp_number")
}
