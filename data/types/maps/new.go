package maps

import (
    "github.com/lucasew/golisp/data"
    "github.com/lucasew/golisp/data/types"
)

type Mappeable interface {
    data.LispValue
    GettableElements() map[string]func()data.LispValue
    SettableElements() map[string]func(data.LispValue)data.LispValue
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
