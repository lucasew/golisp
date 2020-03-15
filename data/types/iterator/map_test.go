package iterator

import (
	"github.com/lucasew/golisp/data"
	"github.com/lucasew/golisp/data/entity/test"
	"github.com/lucasew/golisp/data/types"
	"testing"
)

func justReturnFirst(v ...data.LispValue) (data.LispValue, error) {
	return v[0], nil
}

func TestMapIterator(t *testing.T) {
	v := NewMapIterator(NewRangeIteratorTo(3), types.NewFunction(justReturnFirst))
	test.TestValues(v, t, "lisp_iterator")
	expected := NewRangeIteratorTo(3)
	for i := 0; i < 3; i++ {
		if expected.Next().Repr() != v.Next().Repr() {
			t.Fail()
		}
	}
	if !v.IsEnd() {
		t.Fail()
	}
}
