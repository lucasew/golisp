package iterator

import (
	"github.com/lucasew/golisp/data/entity/test"
	"github.com/lucasew/golisp/data/types"
	"testing"
)

func TestConsIterator(t *testing.T) {
	v := NewConsIterator(types.NewCons(types.Nil))
	test.TestValues(v, t, "lisp_iterator")
}
