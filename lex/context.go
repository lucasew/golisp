package lex

import (
	"context"
	"fmt"
)

type ParseContext struct {
	context.Context
	l *Context
}

func (pc ParseContext) Lex() *Context {
	return pc.l
}

func NewParseContextFromContext(ctx context.Context, l *Context) ParseContext {
	return ParseContext{
		ctx,
		l,
	}
}

type Context struct {
	data  []byte
	index int
}

func NewContext(data []byte) *Context {
	return &Context{
		data:  data,
		index: 0,
	}
}

func (c *Context) MustGetByte() LexByte {
	b, ok := c.GetByte()
	if !ok {
		panic(fmt.Errorf("%w: invalid operation", ErrEOF))
	}
	return b
}

func (c *Context) GetByte() (chr LexByte, ok bool) {
	if c.IsEOF() {
		return 0, false
	}
	return LexByte(c.data[c.index]), true
}

func (c *Context) IsEOF() bool {
	return c.index >= len(c.data)
}

func (c *Context) Increment() {
	c.index++
}

func (c *Context) Decrement() {
	c.index--
}

func (c *Context) Index() int {
	return c.index
}

func (c *Context) Slice(from, to int) string {
	s := c.data[from:to]
	return string(s)
}
