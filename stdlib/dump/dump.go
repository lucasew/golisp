package libdump

import (
	"context"
	"github.com/davecgh/go-spew/spew"
	"github.com/lucasew/golisp/data"
	"github.com/lucasew/golisp/data/types"
	"github.com/lucasew/golisp/utils/enforce"
	"github.com/lucasew/golisp/utils/params"
	"github.com/lucasew/golisp/vm"
)

func init() {
	register("env-dump", EnvDump)
	register("spew", Dump)
	register("dump-call", DumpCall)
}

func EnvDump(ctx context.Context, env vm.LispVM, v ...data.LispValue) (data.LispValue, error) {
	err := enforce.Validate(enforce.Length(v, 0))
	if err != nil {
		return types.Nil, err
	}
	return types.NewString(spew.Sdump(env)), nil
}

func Dump(ctx context.Context, v ...data.LispValue) (data.LispValue, error) {
	err := enforce.Validate(enforce.Length(v, 1))
	if err != nil {
		return types.Nil, err
	}
	return types.NewString(spew.Sdump(v[0])), nil
}

func DumpCall(ctx context.Context, v ...data.LispValue) (data.LispValue, error) {
	p := params.NewParameterLookup(ctx, v...)
	return types.NewString(spew.Sdump(p)), nil
}
