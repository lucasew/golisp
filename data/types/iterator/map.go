package iterator

import (
	"github.com/lucasew/golisp/data"
	"github.com/lucasew/golisp/data/types"
)

type MapIterator struct {
    in data.LispIterator
    f data.LispFunction
}

func NewMapIterator(c data.LispIterator, f data.LispFunction) data.LispIterator {
    return &MapIterator{
        in: c,
        f: f,
    }
}

func (m MapIterator) IsEnd() bool {
    return m.in.IsEnd()
}

func (m MapIterator) IsNil() bool {
    return m.in.IsNil()
}

func (MapIterator) LispTypeName() string {
    return "iterator"
}

func (m *MapIterator) Next() data.LispValue {
    v := m.Next()
    r, err := m.f.LispCall(types.NewCons(v))
    if err != nil {
        return types.Nil
    }
    return r
}

func (MapIterator) Repr() string {
    return "< map iterator >"
}
