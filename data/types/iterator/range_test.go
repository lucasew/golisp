package iterator

import (
	"github.com/lucasew/golisp/data/entity/test"
	"testing"
)

func TestRangeIterator(t *testing.T) {
	v := NewRangeIteratorTo(10)
	test.TestValues(v, t, "lisp_iterator")
}
