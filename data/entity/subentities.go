package entity

import (
	"github.com/lucasew/golisp/data"
)

func NewSubEntities() data.LispSubEntities {
	return SubEntities{
		entities: map[string]data.LispSubEntity{},
	}
}

type SubEntities struct {
	entities map[string]data.LispSubEntity
}

func (su SubEntities) LookupSubEntity(k string) data.LispSubEntity {
	if su.entities == nil {
		return nil
	}
	e, ok := su.entities[k]
	if !ok {
		return nil
	}
	return e
}

func (su SubEntities) Keys() []string {
	ret := []string{}
	if su.entities == nil {
		return ret
	}
	for k := range su.entities {
		ret = append(ret, k)
	}
	return ret
}

func (su SubEntities) AddSubEntity(k string, s data.LispSubEntity) data.LispSubEntities {
	_, ok := su.entities[k]
	if ok {
		return nil
	}
	su.entities[k] = s
	return su
}
