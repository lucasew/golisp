package iterator

import (
	"context"
	"github.com/lucasew/golisp/data"
	"github.com/lucasew/golisp/data/entity/test"
	"github.com/lucasew/golisp/data/types"
	"testing"
)

var ctx = context.Background()

func justReturnFirst(ctx context.Context, v ...data.LispValue) (data.LispValue, error) {
	return v[0], nil
}

func TestMapIterator(t *testing.T) {
	v := NewMapIterator(NewRangeIteratorTo(3), types.NewFunction(justReturnFirst))
	test.TestValues(v, t, "lisp_iterator")
	expected := NewRangeIteratorTo(3)
	for i := 0; i < 3; i++ {
		if expected.Next(ctx).Repr() != v.Next(ctx).Repr() {
			t.Fail()
		}
	}
	if !v.IsEnd(ctx) {
		t.Fail()
	}
}
