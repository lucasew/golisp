package stdlib

import (
	"context"
	"github.com/lucasew/golisp/data"
	"github.com/lucasew/golisp/data/types"
	"github.com/lucasew/golisp/data/types/iterator"
	"github.com/lucasew/golisp/data/types/maps"
	"github.com/lucasew/golisp/utils/enforce"
	"github.com/lucasew/golisp/utils/params"
)

func init() {
	register("new-map", NewMap)
	register("map-set", MapSet)
	register("map-get", MapGet)
	register("map-keys", MapKeys)
	register("map-values", MapValues)
	register("map-tuples", MapTuples)
}

func NewMap(ctx context.Context, v ...data.LispValue) (data.LispValue, error) {
	p := params.NewParameterLookup(ctx, v...)
	err := enforce.Validate(enforce.Length(p.Args, 0))
	if err != nil {
		return types.Nil, err
	}
	return maps.NewMapFromMapString(p.KwArgs), nil
}

func MapSet(ctx context.Context, v ...data.LispValue) (data.LispValue, error) {
	err := enforce.Validate(
		enforce.Length(v, 3),
		enforce.Entity("lisp_namespace", v, 1),
	)
	if err != nil {
		return types.Nil, err
	}
	m := v[0].(data.LispNamespace)
	return m.Set(v[1], v[2]), nil
}

func MapGet(ctx context.Context, v ...data.LispValue) (data.LispValue, error) {
	err := enforce.Validate(
		enforce.Length(v, 2),
		enforce.Entity("lisp_namespace", v, 1),
	)
	if err != nil {
		return types.Nil, err
	}
	m := v[0].(data.LispNamespace)
	return m.Get(v[1]), nil
}

func MapKeys(ctx context.Context, v ...data.LispValue) (data.LispValue, error) {
	err := enforce.Validate(
		enforce.Length(v, 1),
		enforce.Entity("lisp_map", v, 1),
	)
	if err != nil {
		return types.Nil, err
	}
	m := v[0].(data.LispMap)
	return iterator.NewIterator(m.Keys())
}

func MapValues(ctx context.Context, v ...data.LispValue) (data.LispValue, error) {
	err := enforce.Validate(
		enforce.Length(v, 1),
		enforce.Entity("lisp_map", v, 1),
	)
	if err != nil {
		return types.Nil, err
	}
	m := v[0].(data.LispMap)
	return iterator.NewIterator(m.Values())
}

func MapTuples(ctx context.Context, v ...data.LispValue) (data.LispValue, error) {
	err := enforce.Validate(
		enforce.Length(v, 1),
		enforce.Entity("lisp_map", v, 1),
	)
	if err != nil {
		return types.Nil, err
	}
	m := v[0].(data.LispMap)
	return iterator.NewIterator(m.Tuples())
}
