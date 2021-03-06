package types

import (
	"fmt"
	"github.com/lucasew/golisp/data"
	"github.com/lucasew/golisp/data/entity/register"
	"github.com/lucasew/golisp/data/types/number"
)

func init() {
	register.Register("string", func(v data.LispValue) bool {
		_, ok := v.(String)
		return ok
	})
}

type String string

func IsString(v data.LispValue) bool {
	_, ok := v.(data.LispString)
	return ok
}

func NewString(s string) String {
	return String(s)
}

func (c String) UnwrapCons() ([]data.LispValue, error) {
	ret := make([]data.LispValue, len(c))
	for k := range c {
		ret[k] = number.NewByte(c[k])
	}
	return ret, nil
}

func (c String) Get(k int) data.LispValue {
	if k >= c.Len() {
		return Nil
	}
	return number.NewByte(c[k])
}

func (c String) Len() int {
	return len(c)
}

func (c String) IsNil() bool {
	return len(c) == 0
}

func (c String) Car() data.LispValue {
	if len(c) == 0 {
		return Nil
	}
	s := string(c)[0]
	return number.NewByte(s)
}

func (c String) Cdr() data.LispCarCdr {
	if len(c) < 2 {
		return Nil
	}
	s := string(c)[1:len(c)] // TODO: Testar se não teremos acessos errados aqui
	return NewString(s)
}

func (c String) Repr() string {
	return fmt.Sprintf("\"%s\"", c)
}

func (c String) ToString() string {
	return string(c)
}

func (String) LispTypeName() string {
	return "string"
}
