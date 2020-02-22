package iterator

import (
    "testing"
    "github.com/lucasew/golisp/data/types"
)

func TestRangeIterator(t *testing.T) {
    f := NewRangeIteratorTo(10)
    t.Run("iterator", IteratorTest(f))
    t.Run("value", types.ValueTest(f))
}

