package types

import (
    "github.com/lucasew/golisp/data"
)

type Portal chan(data.LispValue)

func NewPortal(buf int) data.LispPortal {
    return Portal(make(chan(data.LispValue), buf))
}

func (Portal) Car() data.LispValue {
    return data.Nil
}

func (Portal) Cdr() data.LispValue {
    return data.Nil
}

func (p Portal) IsNil() bool {
    return p == nil
}

func (p Portal) Repr() string {
    return "<native portal>"
}

func (p Portal) Send(v data.LispValue) data.LispValue {
    p <- v
    return v
}

func (p Portal) SendUnblocking(v data.LispValue) data.LispValue {
    select {
    case p <- v:
        return v
    default:
        return data.Nil
    }
}

func (p Portal) Recv() data.LispValue {
    v := <- p
    return v
}

func (p Portal) RecvUnblocking() data.LispValue {
    select {
    case v := <- p:
        return v
    default:
        return data.Nil
    }
}