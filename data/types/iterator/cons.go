package iterator

import (
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

func (c *ConsIterator) Next() data.LispValue {
	if c.IsEnd() {
		return types.Nil
	}
	v := c.cons.Get(c.index)
	c.index++
	return v
}

func (c ConsIterator) IsEnd() bool {
	return c.cons.Len() <= c.index
}

func (c ConsIterator) IsNil() bool {
	return c.IsEnd()
}

func (ConsIterator) LispTypeName() string {
	return "iterator"
}

func (ConsIterator) Repr() string {
	return "< iterator >"
}
