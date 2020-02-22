package iterator

import (
    "testing"
    "github.com/lucasew/golisp/data/types"
)

func TestConsIterator(t *testing.T) {
    f := NewConsIterator(types.NewCons(types.Nil))
    t.Run("iterator", IteratorTest(f))
    t.Run("value", types.ValueTest(f))
}

var IteratorTest = types.NewTestHelper(IsIterator)
