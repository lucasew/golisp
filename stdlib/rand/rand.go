package rand

import (
	"github.com/lucasew/golisp/data"
	"github.com/lucasew/golisp/data/types"
	"github.com/lucasew/golisp/data/types/number"
	"github.com/lucasew/golisp/utils/enforce"
	"math/rand"
)

func init() {
	register("rand-int", RandInt)
	register("rand-cons", RandCons)
	register("rand-map", RandMap)
}

func RandInt(v ...data.LispValue) (data.LispValue, error) {
	err := enforce.Validate(
		enforce.Length(v, 1),
		enforce.Entity("int", v, 1),
	)
	if err != nil {
		return types.Nil, err
	}
	n, _ := v[0].(number.LispInt).Int64()
	rdn := rand.Int63n(n)
	return number.NewIntFromInt64(rdn), nil
}

func RandCons(v ...data.LispValue) (data.LispValue, error) {
	err := enforce.Validate(
		enforce.Length(v, 1),
		enforce.Entity("lisp_cons", v, 1),
	)
	if err != nil {
		return types.Nil, err
	}
	cons := v[0].(data.LispCons)
	l := cons.Len()
	return cons.Get(rand.Intn(l)), nil
}

func RandMap(v ...data.LispValue) (data.LispValue, error) {
	err := enforce.Validate(
		enforce.Length(v, 1),
		enforce.Entity("lisp_map", v, 1),
	)
	if err != nil {
		return types.Nil, err
	}
	m := v[0].(data.LispMap)
	k, err := RandCons(m.Keys())
	if err != nil {
		// Isto não pode acontecer
		panic(err)
	}
	return types.NewCons(k, m.Get(k)), nil
}
