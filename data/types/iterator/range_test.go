package iterator

import (
	"github.com/lucasew/golisp/data/types"
	"testing"
)

func TestRangeIterator(t *testing.T) {
	f := NewRangeIteratorTo(10)
	t.Run("iterator", IteratorTest(f))
	t.Run("value", types.ValueTest(f))
}
