package maps

import (
	"github.com/lucasew/golisp/data"
	"github.com/lucasew/golisp/data/types"
)

type MapValue map[data.LispValue]data.LispValue

func NewMapFromMapValue(v map[data.LispValue]data.LispValue) data.LispMap {
    return MapValue(v)
}

func (m MapValue) Unwrap() map[data.LispValue]data.LispValue {
    return map[data.LispValue]data.LispValue(m)
}

func (m MapValue) IsNil() bool {
    return len(m.Unwrap()) == 0
}

func (m MapValue) LispTypeName() string {
    return "map"
}

func (m MapValue) Repr() string {
    return "< todo map >"
}

func (m MapValue) Get(k data.LispValue) data.LispValue {
    v, ok := m.Unwrap()[k]
    if !ok {
        return types.Nil
    }
    return v
}

func (m MapValue) Set(k data.LispValue, v data.LispValue) data.LispValue {
    m.Unwrap()[k] = v
    return m
}

func (m MapValue) Keys() data.LispCons {
    ret := make([]data.LispValue, len(m.Unwrap()))
    i := 0
    for k := range m.Unwrap() {
        ret[i] = k
        i++
    }
    return types.NewCons(ret...)
}

func (m MapValue) Values() data.LispCons {
    ret := make([]data.LispValue, len(m.Unwrap()))
    i := 0
    for _, v := range m.Unwrap() {
        ret[i] = v
        i++
    }
    return types.NewCons(ret...)
}

func (m MapValue) Tuples() data.LispCons {
    ret := make([]data.LispValue, len(m.Unwrap()))
    i := 0
    for k, v := range m.Unwrap() {
        ret[i] = types.NewCons(k, v)
        i++
    }
    return types.NewCons(ret...)
}

func (m MapValue) Len() int {
    return len(m.Unwrap())
}
