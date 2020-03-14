package stdlib

import (
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

func NewMap(v ...data.LispValue) (data.LispValue, error) {
    p := params.NewParameterLookup(v...)
	err := enforce.Validate(enforce.Length(p.Args, 0))
	if err != nil {
		return types.Nil, err
	}
	return maps.NewMapFromMapString(p.KwArgs), nil
}

func MapSet(v ...data.LispValue) (data.LispValue, error) {
	err := enforce.Validate(
		enforce.Length(v, 3),
		enforce.Namespace(v, 1),
	)
	if err != nil {
		return types.Nil, err
	}
	m := v[0].(data.LispNamespace)
	return m.Set(v[1], v[2]), nil
}

func MapGet(v ...data.LispValue) (data.LispValue, error) {
	err := enforce.Validate(
		enforce.Length(v, 2),
		enforce.Namespace(v, 1),
	)
	if err != nil {
		return types.Nil, err
	}
	m := v[0].(data.LispNamespace)
	return m.Get(v[1]), nil
}

func MapKeys(v ...data.LispValue) (data.LispValue, error) {
	err := enforce.Validate(
		enforce.Length(v, 1),
		enforce.Map(v, 1),
	)
	if err != nil {
		return types.Nil, err
	}
	m := v[0].(data.LispMap)
	return iterator.NewIterator(m.Keys())
}

func MapValues(v ...data.LispValue) (data.LispValue, error) {
	err := enforce.Validate(
		enforce.Length(v, 1),
		enforce.Map(v, 1),
	)
	if err != nil {
		return types.Nil, err
	}
	m := v[0].(data.LispMap)
	return iterator.NewIterator(m.Values())
}

func MapTuples(v ...data.LispValue) (data.LispValue, error) {
	err := enforce.Validate(
		enforce.Length(v, 1),
		enforce.Map(v, 1),
	)
	if err != nil {
		return types.Nil, err
	}
	m := v[0].(data.LispMap)
	return iterator.NewIterator(m.Tuples())
}
