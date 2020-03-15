package iterator

import (
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

func NewFilterIterator(c data.LispIterator, f data.LispFunction) data.LispIterator {
	r := &FilterIterator{
		in:  c,
		f:   f,
		end: false,
	}
	r.next = r.nextMatched()
	return r
}

func (m *FilterIterator) nextMatched() data.LispValue {
begin:
	if m.in.IsEnd() {
		m.end = true
		return types.Nil
	}
	v := m.in.Next()
	r, err := m.f.LispCall(v)
	if err != nil {
		return types.Nil
	}
	if r.IsNil() {
		goto begin
	}
	return v
}

func (m FilterIterator) IsEnd() bool {
	return m.end
}

func (m FilterIterator) IsNil() bool {
	return m.IsEnd()
}

func (FilterIterator) LispTypeName() string {
	return "iterator"
}

func (m *FilterIterator) Next() data.LispValue {
	r := m.next
	m.next = m.nextMatched()
	return r
}

func (FilterIterator) Repr() string {
	return "< filter iterator >"
}
