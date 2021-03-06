package types

import (
	"github.com/lucasew/golisp/data"
	"github.com/lucasew/golisp/data/entity/register"
)

func init() {
	register.Register("portal", func(v data.LispValue) bool {
		_, ok := v.(Portal)
		return ok
	})
}

type Portal chan (data.LispValue)

func IsPortal(v data.LispValue) bool {
	_, ok := v.(data.LispPortal)
	return ok
}

func NewPortal(buf int) data.LispPortal {
	return Portal(make(chan (data.LispValue), buf))
}

func (p Portal) IsNil() bool {
	return p == nil
}

func (p Portal) Repr() string {
	return "<native portal>"
}

func (p Portal) Close() {
	close(p)
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
		return Nil
	}
}

func (p Portal) Recv() data.LispValue {
	v := <-p
	return v
}

func (p Portal) RecvUnblocking() data.LispValue {
	select {
	case v := <-p:
		return v
	default:
		return Nil
	}
}

func (Portal) LispTypeName() string {
	return "portal"
}
