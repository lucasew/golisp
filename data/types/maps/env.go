package maps

import (
	"github.com/lucasew/golisp/data"
	"github.com/lucasew/golisp/data/types"
	"github.com/lucasew/golisp/vm/components/env"
)

type LispEnv struct {
	env.LispEnv
}

func NewLispNamespaceFromEnv(e env.LispEnv) data.LispNamespace {
	return LispEnv{e}
}

func (e LispEnv) Get(k data.LispValue) data.LispValue {
	ks, ok := k.(data.LispString)
	if !ok {
		return types.Nil
	}
	return e.Get(ks)
}

func (e LispEnv) Set(k data.LispValue, v data.LispValue) data.LispValue {
	ks, ok := k.(data.LispString)
	if !ok {
		return types.Nil
	}
	return e.SetLocal(ks.ToString(), v)
}

func (e LispEnv) IsNil() bool {
	return e.IsNil()
}

func (e LispEnv) LispTypeName() string {
	return "map"
}

func (e LispEnv) Repr() string {
	return "< lisp env >"
}
