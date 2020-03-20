package entity

import (
	"github.com/lucasew/golisp/data"
)

func NewSubEntity(
	get func(data.LispValue) data.LispValue,
	set func(self, v data.LispValue) data.LispValue,
) SubEntity {
	return SubEntity{get, set}
}

type SubEntity struct {
	get func(v data.LispValue) data.LispValue
	set func(v data.LispValue, to data.LispValue) data.LispValue
}

func (s SubEntity) Get(self data.LispValue) data.LispValue {
	return s.get(self)
}

func (s SubEntity) Set(self, to data.LispValue) data.LispValue {
	return s.set(self, to)
}
