package stdlib

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/lucasew/golisp/data"
	"github.com/lucasew/golisp/data/types"
	"github.com/lucasew/golisp/utils/enforce"
	"github.com/lucasew/golisp/vm"
)

func init() {
	register("env-dump", EnvDump)
}

func EnvDump(env vm.LispVM, v ...data.LispValue) (data.LispValue, error) {
	err := enforce.Validate(enforce.Length(v, 0))
	if err != nil {
		return types.Nil, err
	}
	return types.NewString(spew.Sdump(env)), nil
}
