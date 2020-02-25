package iterator

import (
	"github.com/lucasew/golisp/data"
	"github.com/lucasew/golisp/data/types"
)

type FilterIterator struct {
    in data.LispIterator
    f data.LispFunction
}

func NewFilterIterator(c data.LispIterator, f data.LispFunction) data.LispIterator {
    return &FilterIterator{
        in: c,
        f: f,
    }
}

func (m FilterIterator) IsEnd() bool {
    return m.in.IsEnd()
}

func (m FilterIterator) IsNil() bool {
    return m.in.IsNil()
}

func (FilterIterator) LispTypeName() string {
    return "iterator"
}

func (m *FilterIterator) Next() data.LispValue {
    for !m.IsEnd() {
        v := m.Next()
        r, err := m.f.LispCall(types.NewCons(v))
        if err != nil {
            return types.Nil
        }
        if r.IsNil() {
            continue
        }
        return r
    }
    return types.Nil
}

func (FilterIterator) Repr() string {
    return "< filter iterator >"
}
