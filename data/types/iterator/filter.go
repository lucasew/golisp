package iterator

import (
	"context"
	"github.com/lucasew/golisp/data"
	"github.com/lucasew/golisp/data/entity/register"
	"github.com/lucasew/golisp/data/types"
)

func init() {
	register.Register("filter_iterator", func(v data.LispValue) bool {
		_, ok := v.(FilterIterator)
		return ok
	})
}

type FilterIterator struct {
	in   data.LispIterator
	f    data.LispFunction
	next data.LispValue
	end  bool
}

func NewFilterIterator(ctx context.Context, c data.LispIterator, f data.LispFunction) data.LispIterator {
	r := &FilterIterator{
		in:  c,
		f:   f,
		end: false,
	}
	r.next = r.nextMatched(ctx)
	return r
}

func (m *FilterIterator) nextMatched(ctx context.Context) data.LispValue {
begin:
	if m.in.IsEnd(ctx) {
		m.end = true
		return types.Nil
	}
	v := m.in.Next(ctx)
	r, err := m.f.LispCall(ctx, v)
	if err != nil {
		return types.Nil
	}
	if r.IsNil() {
		goto begin
	}
	return v
}

func (m FilterIterator) IsEnd(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true
	default:
		return m.end
	}
}

func (m FilterIterator) IsNil() bool {
	return m.IsEnd(context.TODO())
}

func (FilterIterator) LispTypeName() string {
	return "iterator"
}

func (m *FilterIterator) Next(ctx context.Context) data.LispValue {
	r := m.next
	m.next = m.nextMatched(ctx)
	return r
}

func (FilterIterator) Repr() string {
	return "< filter iterator >"
}
