package params

import (
    "github.com/lucasew/golisp/data"
    "github.com/lucasew/golisp/data/types/iterator"
    "github.com/lucasew/golisp/data/types"
    "fmt"
)

func NewParameterLookup(v ...data.LispValue) ParameterLookup {
    params := ParameterLookup{
        []data.LispValue{}, 
        map[string]data.LispValue{},
    }
    values := iterator.NewConsIterator(types.NewCons(v...))
    for !values.IsEnd() {
        switch this := values.Next().(type) {
        case types.Atom:
            params.KwArgs[this.AtomString()] = values.Next()
        default:
            params.Args = append(params.Args, this)
        }
    }
    return params
}

type ParameterLookup struct {
    Args []data.LispValue
    KwArgs map[string]data.LispValue
}

func (p ParameterLookup) GetNth(i int, criteria func(data.LispValue) error) (data.LispValue, error) {
    if len(p.Args) < i {
        return types.Nil, fmt.Errorf("not enough arguments. requested %dith got length %d", i, len(p.Args))
    }
    data := p.Args[i]
    if criteria == nil {
        return data, nil
    }
    err := criteria(data)
    if err != nil {
        return types.Nil, fmt.Errorf("getNth:%d %w", i, err)
    }
    return data, nil
}

func (p ParameterLookup) GetKey(k string, criteria func(data.LispValue)error) (data.LispValue, error) {
    v, ok := p.KwArgs[k]
    if !ok {
        return types.Nil, fmt.Errorf("undefined key: %s", k)
    }
    if criteria == nil {
        return v, nil
    }
    err := criteria(v)
    if err != nil {
        return types.Nil, fmt.Errorf("getKey:%s %w", k, err)
    }
    return v, nil
}
