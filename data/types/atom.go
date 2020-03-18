package types

import (
	"fmt"
	"github.com/lucasew/golisp/data"
	"github.com/lucasew/golisp/data/entity"
	"github.com/lucasew/golisp/data/entity/register"
	_ "github.com/lucasew/golisp/data/types/test"
)

func init() {
	register.Register(new(Atom).LispEntity())
}

func (Atom) LispEntity() data.LispEntity {
	return entity.Entity{
		"atom", func(v data.LispValue) bool {
			_, ok := v.(Atom)
			return ok
		},
	}
}

type Atom string

func NewAtom(s string) data.LispAtom {
	return Atom(s)
}

func (a Atom) IsNil() bool {
	return string(a) == "nil"
}

func (a Atom) Repr() string {
	return fmt.Sprintf(":%s", a)
}

func (Atom) LispTypeName() string {
	return "atom"
}

func (a Atom) AtomString() string {
	return string(a)
}
