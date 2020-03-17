package number

import (
	"github.com/lucasew/golisp/data/entity/test"
	_ "github.com/lucasew/golisp/data/types/test"
	"testing"
)

func TestByteType(t *testing.T) {
	v := NewByte(0)
	test.TestValues(v, t, "byte", "lisp_number")
}
