package env

import (
    "github.com/lucasew/golisp/data"
)

type LispEnv struct {
    upper *LispEnv
    data map[string]data.LispValue
}

func NewLispEnv(upper *LispEnv, kv ...LispEnvKV) *LispEnv {
    e := LispEnv{
        upper: upper,
        data: map[string]data.LispValue{},
    }
    for _, v := range kv {
        e.SetLocal(v.Key, v.Value)
    }
    return &e
}

func (e *LispEnv) SetLocal(key string, value data.LispValue) data.LispValue {
    e.data[key] = value
    return value
}

func (e *LispEnv) SetGlobal(key string, value data.LispValue) data.LispValue {
    if e.upper == nil {
        e.data[key] = value
        return value
    } else {
        return e.upper.SetGlobal(key, value)
    }
}

func (e *LispEnv) Get(key string) data.LispValue {
    r, ok := e.data[key]
    if !ok {
        if e.upper != nil {
            return e.upper.Get(key)
        }
    } else {
        return r
    }
    return data.Nil
}


