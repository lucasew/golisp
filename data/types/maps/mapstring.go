package maps

import (
	"github.com/lucasew/golisp/data"
	"github.com/lucasew/golisp/data/types"
)

type MapString map[string]data.LispValue

func NewMapFromMapString(v map[string]data.LispValue) data.LispMap {
    return MapString(v)
}

func (m MapString) Unwrap() map[string]data.LispValue {
    return map[string]data.LispValue(m)
}

func (m MapString) IsNil() bool {
    return len(m.Unwrap()) == 0
}

func (m MapString) LispTypeName() string {
    return "map"
}

func (m MapString) Repr() string {
    return "< todo map >"
}

func (m MapString) Get(k data.LispValue) data.LispValue {
    ks, ok := k.(data.LispString)
    if !ok {
        return types.Nil
    }
    v, ok := m.Unwrap()[ks.ToString()]
    if !ok {
        return types.Nil
    }
    return v
}

func (m MapString) Set(k data.LispValue, v data.LispValue) data.LispValue {
    ks, ok := k.(data.LispString)
    if !ok {
        return types.Nil
    }
    m.Unwrap()[ks.ToString()] = v
    return m
}

func (m MapString) Keys() data.LispCons {
    ret := make([]data.LispValue, len(m.Unwrap()))
    i := 0
    for k := range m.Unwrap() {
        ret[i] = types.NewString(k)
        i++
    }
    return types.NewCons(ret...)
}

func (m MapString) Values() data.LispCons {
    ret := make([]data.LispValue, len(m.Unwrap()))
    i := 0
    for _, v := range m.Unwrap() {
        ret[i] = v
        i++
    }
    return types.NewCons(ret...)
}

func (m MapString) Tuples() data.LispCons {
    ret := make([]data.LispValue, len(m.Unwrap()))
    i := 0
    for k, v := range m.Unwrap() {
        ret[i] = types.NewCons(types.NewString(k), v)
        i++
    }
    return types.NewCons(ret...)
}




func (m MapString) Len() int {
    return len(m.Unwrap())
}
