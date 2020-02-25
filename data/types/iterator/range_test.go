package iterator

import (
	"github.com/lucasew/golisp/data/types/test"
	"testing"
)

func TestRangeIterator(t *testing.T) {
	v := NewRangeIteratorTo(10)
	t.Run("value", test.NewTestHelper(test.IsValue)(v))
	t.Run("iterator", test.NewTestHelper(test.IsIterator)(v))
}
