package types

import (
	"github.com/lucasew/golisp/data"
	"github.com/lucasew/golisp/data/entity"
	"github.com/lucasew/golisp/data/entity/register"
	"github.com/lucasew/golisp/data/types/number"
	_ "github.com/lucasew/golisp/data/types/test"
)

func init() {
	register.Register(new(Symbol).LispEntity())
}

func (Symbol) LispEntity() data.LispEntity {
	return entity.Entity{
		"symbol", func(v data.LispValue) bool {
			_, ok := v.(Symbol)
			return ok
		},
	}
}

type Symbol string

func IsSymbol(v data.LispValue) bool {
	_, ok := v.(Symbol)
	return ok
}

func NewSymbol(s string) Symbol {
	return Symbol(s)
}

func (s Symbol) UnwrapCons() ([]data.LispValue, error) {
	return String(s).UnwrapCons()
}

func (s Symbol) Get(k int) data.LispValue {
	return String(k).Get(k)
}

func (s Symbol) Len() int {
	return len(s)
}

func (s Symbol) ToString() string {
	return string(s)
}

func (Symbol) IsNil() bool {
	return false
}

func (s Symbol) Repr() string {
	return string(s)
}

func (Symbol) LispTypeName() string {
	return "symbol"
}

func (s Symbol) Car() data.LispValue {
	return number.NewByte(s[0])
}

func (s Symbol) Cdr() data.LispCarCdr {
	return NewSymbol(String(s).Cdr().(data.LispString).ToString())
}
