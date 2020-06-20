package iterator

import (
	"context"
	"github.com/lucasew/golisp/data"
	"github.com/lucasew/golisp/data/entity/register"
	"github.com/lucasew/golisp/data/types"
	"github.com/lucasew/golisp/data/types/number"
)

func init() {
	register.Register("range_iterator", func(v data.LispValue) bool {
		_, ok := v.(RangeIterator)
		return ok
	})
}

type RangeIterator struct {
	state int
	to    int
	step  int
}

func NewRangeIteratorTo(i int) data.LispIterator {
	return &RangeIterator{
		state: 0,
		to:    i,
		step:  1,
	}
}

func (r RangeIterator) IsEnd(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true
	default:
		return r.state >= r.to
	}
}

func (r RangeIterator) IsNil() bool {
	return r.IsEnd(context.TODO())
}

func (r RangeIterator) LispTypeName() string {
	return "iterator"
}

func (r *RangeIterator) Next(ctx context.Context) data.LispValue {
	v := r.state
	r.state += r.step
	if r.state > r.to {
		return types.Nil
	}
	return number.NewIntFromInt64(int64(v))
}

func (RangeIterator) Repr() string {
	return "< iterator >"
}
