package types

import (
	"github.com/lucasew/golisp/data/entity/test"
	"testing"
)

func TestAtomType(t *testing.T) {
	v := NewAtom("")
	test.TestValues(v, t, "atom")
}
