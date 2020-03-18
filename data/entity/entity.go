package entity

import (
	"fmt"
	"github.com/lucasew/golisp/data"
	"reflect"
)

type Entity struct {
	Name string
	Isfn func(data.LispValue) bool
}

func (e Entity) EntityName() string {
	return e.Name
}

func (e Entity) EntityIsFn(v data.LispValue) bool {
	if e.Isfn == nil {
		return false
	}
	return e.Isfn(v)
}

func NewSameTypeEntity(a data.LispValue) Entity {
	ta := reflect.TypeOf(a)
	return Entity{
		Name: fmt.Sprintf("same.%s", a.LispEntity().EntityName()),
		Isfn: func(v data.LispValue) bool {
			return reflect.TypeOf(v).AssignableTo(ta)
		},
	}
}
