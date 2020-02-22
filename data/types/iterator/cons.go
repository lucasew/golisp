package iterator

import (
    "github.com/lucasew/golisp/data"
    "github.com/lucasew/golisp/data/types"
)

type ConsIterator struct {
    cons data.LispCons
    index int
}

func NewConsIterator(c data.LispCons) data.LispIterator {
    return &ConsIterator{
        cons: c,
        index: 0,
    }
}

func (c *ConsIterator) Next() data.LispValue {
    v := c.cons.Get(c.index)
    if c.IsEnd() {
        return types.Nil
    }
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
