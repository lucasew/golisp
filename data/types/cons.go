package types

import (
	"fmt"
	"github.com/lucasew/golisp/data"
	"github.com/lucasew/golisp/data/entity"
	"github.com/lucasew/golisp/data/entity/register"
	"strings"
)

func init() {
	register.Register(new(Cons).LispEntity())
}

func (Cons) LispEntity() data.LispEntity {
	return entity.Entity{
		"cons", func(v data.LispValue) bool {
			_, ok := v.(Cons)
			return ok
		},
	}
}

type Cons []data.LispValue

func NewCons(v ...data.LispValue) data.LispCons {
	return Cons(v)
}

func (i Cons) UnwrapCons() ([]data.LispValue, error) {
	return []data.LispValue(i), nil
}

func (i Cons) Car() data.LispValue {
	if len(i) == 0 {
		return Nil
	}
	return i[0]
}

func (i Cons) Cdr() data.LispCarCdr {
	if len(i) < 2 {
		return NewCons()
	}
	return i[1:len(i)]
}

func (i Cons) IsNil() bool {
	return len(i) == 0
}

func (i Cons) Repr() string {
	strs := make([]string, len(i))
	for k, v := range i {
		strs[k] = v.Repr()
	}
	return fmt.Sprintf(" (%s) ", strings.Join(strs, " "))
}

func (i Cons) Get(k int) data.LispValue {
	if k > len(i) {
		return Nil
	}
	return i[k]
}

func (i Cons) Len() int {
	return len(i)
}

func (Cons) LispTypeName() string {
	return "cons"
}
