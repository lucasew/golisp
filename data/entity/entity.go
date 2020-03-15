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

func NewSameTypeEntity(a data.LispValue) Entity {
	ta := reflect.TypeOf(a)
	return Entity{
		Name: fmt.Sprintf("same.%s", a.LispTypeName()),
		Isfn: func(v data.LispValue) bool {
			return reflect.TypeOf(v).AssignableTo(ta)
		},
	}
}
