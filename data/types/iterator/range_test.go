package iterator

import (
    "testing"
    "github.com/lucasew/golisp/data/types"
    "fmt"
)

func TestRangeIterator(t *testing.T) {
    f := NewRangeIteratorTo(10)
    t.Run("iterator", IteratorTest(f))
    t.Run("value", types.ValueTest(f))
    v := make([]int, 10)
    for n := f.Next(); !f.IsEnd(); n = f.Next() {
        num, _ := n.(types.LispInt).Int64()
        v[num] = int(num)
    }
    if !(v[0] == 0 && v[5] == 5 && v[9] == 9) {
        fmt.Printf("%+v\n", v)
        t.Fail()
    }
}

