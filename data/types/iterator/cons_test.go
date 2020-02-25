package iterator

import (
	"github.com/lucasew/golisp/data/types/test"
	"github.com/lucasew/golisp/data/types"
	"testing"
)

func TestConsIterator(t *testing.T) {
	v := NewConsIterator(types.NewCons(types.Nil))
	t.Run("value", test.NewTestHelper(test.IsValue)(v))
	t.Run("iterator", test.NewTestHelper(test.IsIterator)(v))
}
