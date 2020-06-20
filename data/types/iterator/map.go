package iterator

import (
	"context"
	"github.com/lucasew/golisp/data"
	"github.com/lucasew/golisp/data/entity"
	"github.com/lucasew/golisp/data/entity/register"
	"github.com/lucasew/golisp/data/types"
)

func init() {
	register.Register(new(MapIterator).LispEntity())
}

func (MapIterator) LispEntity() data.LispEntity {
	return entity.Entity{
		"map_iterator", func(v data.LispValue) bool {
			_, ok := v.(MapIterator)
			return ok
		},
	}
}

type MapIterator struct {
	in data.LispIterator
	f  data.LispFunction
}

func NewMapIterator(c data.LispIterator, f data.LispFunction) data.LispIterator {
	return &MapIterator{
		in: c,
		f:  f,
	}
}

func (m MapIterator) IsEnd(ctx context.Context) bool {
	return m.in.IsEnd(ctx)
}

func (m MapIterator) IsNil() bool {
	return m.in.IsNil()
}

func (MapIterator) LispTypeName() string {
	return "iterator"
}

func (m *MapIterator) Next(ctx context.Context) data.LispValue {
	if m.in.IsEnd(ctx) {
		return types.Nil
	}
	v := m.in.Next(ctx)
	r, err := m.f.LispCall(ctx, v)
	if err != nil {
		return types.Nil
	}
	return r
}

func (MapIterator) Repr() string {
	return "< map iterator >"
}
