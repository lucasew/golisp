package lex

type Context struct {
    data []byte
    index int
}

func (c *Context) MustGetByte() LexByte {
    b, ok := c.GetByte()
    if !ok {
        panic("invalid operation in EOF")
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

func (c *Context) Index() int {
    return c.index
}
