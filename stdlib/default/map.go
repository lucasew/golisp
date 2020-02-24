package stdlib

import (
	"github.com/lucasew/golisp/data"
	"github.com/lucasew/golisp/data/types"
	"github.com/lucasew/golisp/stdlib/default/enforce"
)

func init() {
	register("new-map", NewMap)
	register("map-set", MapSet)
	register("map-get", MapGet)
	register("map-keys", MapKeys)
	register("map-values", MapValues)
	register("map-tuples", MapTuples)
}

func NewMap(v data.LispCons) (data.LispValue, error) {
	err := enforce.Length(v, 0)
	if err != nil {
		return types.Nil, err
	}
	return types.NewMap(), nil
}

func MapSet(v data.LispCons) (data.LispValue, error) {
	err := enforce.Validate(
		enforce.Length(v, 3),
		enforce.Map(v.Car(), 1),
	)
	if err != nil {
		return types.Nil, err
	}
	m := v.Car().(data.LispMap)
	return m.Set(v.Cdr().Car(), v.Cdr().Cdr().Car()), nil
}

func MapGet(v data.LispCons) (data.LispValue, error) {
	err := enforce.Validate(
		enforce.Length(v, 2),
		enforce.Map(v.Car(), 1),
	)
	if err != nil {
		return types.Nil, err
	}
	m := v.Car().(data.LispMap)
	return m.Get(v.Cdr().Car()), nil
}

func MapKeys(v data.LispCons) (data.LispValue, error) {
	err := enforce.Validate(
		enforce.Length(v, 1),
		enforce.Map(v.Car(), 1),
	)
	if err != nil {
		return types.Nil, err
	}
	m := v.Car().(data.LispMap)
	return m.Keys(), nil
}

func MapValues(v data.LispCons) (data.LispValue, error) {
	err := enforce.Validate(
		enforce.Length(v, 1),
		enforce.Map(v.Car(), 1),
	)
	if err != nil {
		return types.Nil, err
	}
	m := v.Car().(data.LispMap)
	return m.Values(), nil
}

func MapTuples(v data.LispCons) (data.LispValue, error) {
	err := enforce.Validate(
		enforce.Length(v, 1),
		enforce.Map(v.Car(), 1),
	)
	if err != nil {
		return types.Nil, err
	}
	m := v.Car().(data.LispMap)
	return m.Tuples(), nil
}
