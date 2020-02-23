package stdlib

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/lucasew/golisp/data"
	"github.com/lucasew/golisp/data/types"
	"github.com/lucasew/golisp/vm"
	"github.com/lucasew/golisp/vm/default/stdlib/enforce"
)

func init() {
	register("env-dump", EnvDump)
}

func EnvDump(env vm.LispVM, v data.LispCons) (data.LispValue, error) {
	err := enforce.Length(v, 0)
	if err != nil {
		return types.Nil, err
	}
	return types.NewString(spew.Sdump(env)), nil
}
