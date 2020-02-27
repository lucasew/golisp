package iterator

import (
	"github.com/lucasew/golisp/data"
	"github.com/lucasew/golisp/data/types"
	"github.com/lucasew/golisp/data/types/test"
	"testing"
)

func justReturnFirst(v data.LispCons) (data.LispValue, error) {
	return v.Car(), nil
}

func TestMapIterator(t *testing.T) {
	v := NewMapIterator(NewRangeIteratorTo(3), types.NewFunction(justReturnFirst))
	t.Run("value", test.NewTestHelper(test.IsValue)(v))
	t.Run("iterator", test.NewTestHelper(test.IsIterator)(v))
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
