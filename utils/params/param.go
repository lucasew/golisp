package params

import (
	"context"
	"fmt"
	"github.com/lucasew/golisp/data"
	"github.com/lucasew/golisp/data/entity"
	eerr "github.com/lucasew/golisp/data/entity/error"
	"github.com/lucasew/golisp/data/types"
	"github.com/lucasew/golisp/data/types/iterator"
)

func NewParameterLookup(ctx context.Context, v ...data.LispValue) ParameterLookup {
	params := ParameterLookup{
		[]data.LispValue{},
		map[string]data.LispValue{},
	}
	values := iterator.NewConsIterator(types.NewCons(v...))
	for !values.IsEnd(ctx) {
		switch this := values.Next(ctx).(type) {
		case types.Atom:
			params.KwArgs[this.AtomString()] = values.Next(ctx)
		default:
			params.Args = append(params.Args, this)
		}
	}
	return params
}

type ParameterLookup struct {
	Args   []data.LispValue
	KwArgs map[string]data.LispValue
}

func (p ParameterLookup) GetNth(i int, entity *entity.Entity) (data.LispValue, error) {
	if len(p.Args) < i {
		return types.Nil, fmt.Errorf("not enough arguments. requested %dith got length %d", i, len(p.Args))
	}
	data := p.Args[i]
	if entity == nil {
		return data, nil
	}
	err := eerr.EntityAsError(*entity)(data)
	if err != nil {
		return types.Nil, fmt.Errorf("getNth:%d %w", i, err)
	}
	return data, nil
}

func (p ParameterLookup) GetKey(k string, entity *entity.Entity) (data.LispValue, error) {
	v, ok := p.KwArgs[k]
	if !ok {
		return types.Nil, fmt.Errorf("undefined key: %s", k)
	}
	if entity == nil {
		return v, nil
	}
	err := eerr.EntityAsError(*entity)(v)
	if err != nil {
		return types.Nil, fmt.Errorf("getKey:%s %w", k, err)
	}
	return v, nil
}
