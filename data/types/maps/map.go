package maps

import (
	"github.com/lucasew/golisp/data"
	"github.com/lucasew/golisp/data/types"
)

type LispMap map[data.LispValue]data.LispValue

func NewMap() data.LispMap {
	return LispMap(map[data.LispValue]data.LispValue{})
}

func (m LispMap) Get(k data.LispValue) data.LispValue {
	r, ok := m[k]
	if !ok {
		return types.Nil
	}
	return r
}

func (m LispMap) Set(k data.LispValue, v data.LispValue) data.LispValue {
	m[k] = v
	return v
}

func (m LispMap) Keys() data.LispCons {
	ret := make([]data.LispValue, len(m))
	i := 0
	for k := range m {
		ret[i] = k
		i++
	}
	return types.Cons(ret)
}

func (m LispMap) Values() data.LispCons {
	ret := make([]data.LispValue, len(m))
	i := 0
	for _, v := range m {
		ret[i] = v
		i++
	}
	return types.Cons(ret)
}

func (m LispMap) Tuples() data.LispCons {
	ret := make([]data.LispValue, len(m))
	i := 0
	for k, v := range m {
		ret[i] = types.Cons([]data.LispValue{k, v})
		i++
	}
	return types.Cons(ret)
}

func (m LispMap) Len() int {
	return len(m)
}

func (m LispMap) LispTypeName() string {
	return "map"
}

func (m LispMap) Repr() string {
	// TODO: Repr
	return "< todo: map >"
}

func (m LispMap) IsNil() bool {
	return len(m) == 0
}
