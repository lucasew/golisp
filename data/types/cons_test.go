package types

import (
	"github.com/lucasew/golisp/data/entity/test"
	"testing"
)

func TestConsType(t *testing.T) {
	v := NewCons()
	test.TestValues(v, t, "lisp_cons", "lisp_carcdr", "cons")
}
