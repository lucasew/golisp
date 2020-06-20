package maps

import (
	"github.com/lucasew/golisp/data"
	"github.com/lucasew/golisp/data/entity/register"
	"github.com/lucasew/golisp/data/types"
)

func init() {
	register.Register("map_mappeable", func(v data.LispValue) bool {
		_, ok := v.(MapMappeable)
		return ok
	})
}

type MapMappeable struct {
	m Mappeable
}

func NewMapFromMappeable(m Mappeable) data.LispMap {
	return MapMappeable{m}
}

func (m MapMappeable) Unwrap() Mappeable {
	return m.m
}

func (m MapMappeable) Get(k data.LispValue) data.LispValue {
	return MappeableGet(m.m, k)
}

func (m MapMappeable) IsNil() bool {
	return m.Len() == 0
}

func (m MapMappeable) Set(k data.LispValue, v data.LispValue) data.LispValue {
	return MappeableSet(m.m, k, v)
}

func (m MapMappeable) Keys() data.LispCons {
	return MappeableKeys(m.m)
}

func (m MapMappeable) Values() data.LispCons {
	return MappeableValues(m.m)
}

func (m MapMappeable) Tuples() data.LispCons {
	return MappeableTuples(m.m)
}

func (m MapMappeable) Len() int {
	return MappeableLen(m.m)
}

func (m MapMappeable) LispTypeName() string {
	return "map"
}

func (m MapMappeable) Repr() string {
	return "< map mappeable >"
}

type Mappeable interface {
	GettableElements() map[string]func() data.LispValue
	SettableElements() map[string]func(data.LispValue) data.LispValue
}

func MappeableGet(m Mappeable, k data.LispValue) data.LispValue {
	ks, ok := k.(data.LispString)
	if !ok {
		return types.Nil
	}
	elem, ok := m.GettableElements()[ks.ToString()]
	if !ok {
		return types.Nil
	}
	return elem()
}

func MappeableSet(m Mappeable, k data.LispValue, v data.LispValue) data.LispValue {
	ks, ok := k.(data.LispString)
	if !ok {
		return types.Nil
	}
	elem, ok := m.SettableElements()[ks.ToString()]
	if !ok {
		return types.Nil
	}
	return elem(v)
}

func MappeableKeys(m Mappeable) data.LispCons {
	keys := []data.LispValue{}
	for k := range m.GettableElements() {
		keys = append(keys, types.NewString(k))
	}
	return types.NewCons(keys...)
}

func MappeableValues(m Mappeable) data.LispCons {
	values := []data.LispValue{}
	for _, v := range m.GettableElements() {
		values = append(values, v())
	}
	return types.NewCons(values...)
}

func MappeableTuples(m Mappeable) data.LispCons {
	kv := []data.LispValue{}
	for k, v := range m.GettableElements() {
		kv = append(kv, types.NewCons(types.NewString(k), v()))
	}
	return types.NewCons(kv...)
}

func MappeableLen(m Mappeable) int {
	// TODO: Colocar settable em consideração
	return len(m.GettableElements())
}
