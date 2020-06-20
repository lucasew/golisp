package iterator

import (
	"context"
	"github.com/lucasew/golisp/data"
	"github.com/lucasew/golisp/data/entity/register"
	"github.com/lucasew/golisp/data/types"
)

func init() {
	register.Register("cons_iterator", func(v data.LispValue) bool {
		_, ok := v.(ConsIterator)
		return ok
	})
}

type ConsIterator struct {
	cons  data.LispCons
	index int
}

func NewConsIterator(c data.LispCons) data.LispIterator {
	return &ConsIterator{
		cons:  c,
		index: 0,
	}
}

func (c *ConsIterator) Next(ctx context.Context) data.LispValue {
	if c.IsEnd(ctx) {
		return types.Nil
	}
	v := c.cons.Get(c.index)
	c.index++
	return v
}

func (c ConsIterator) IsEnd(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true
	default:
		return c.cons.Len() <= c.index
	}
}

func (c ConsIterator) IsNil() bool {
	return c.IsEnd(context.TODO()) // TODO: Test possíveis falhas de segurança
}

func (ConsIterator) LispTypeName() string {
	return "iterator"
}

func (ConsIterator) Repr() string {
	return "< iterator >"
}
