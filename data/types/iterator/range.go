package iterator

import (
    "github.com/lucasew/golisp/data"
    "github.com/lucasew/golisp/data/types"
)

type RangeIterator struct {
    state int
    to int
    step int
}

func NewRangeIteratorTo(i int) data.LispIterator {
    return &RangeIterator{
        state: 0,
        to: i,
        step: 1,
    }
}

func (r RangeIterator) IsEnd() bool {
    return r.state > r.to
}

func (r RangeIterator) IsNil() bool {
    return r.IsEnd()
}

func (r RangeIterator) LispTypeName() string {
    return "iterator"
}

func (r *RangeIterator) Next() data.LispValue {
    v := r.state
    r.state += r.step
    if r.state > r.to {
        return types.Nil
    }
    return types.NewIntFromInt64(int64(v))
}

func (RangeIterator) Repr() string {
    return "< iterator >"
}
