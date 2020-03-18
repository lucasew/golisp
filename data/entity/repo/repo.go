package repo

import (
	"context"
	"fmt"
	"github.com/lucasew/golisp/data"
	"github.com/lucasew/golisp/data/entity/register"
	"github.com/lucasew/golisp/data/types"
	"github.com/lucasew/golisp/stdlib"
	"github.com/lucasew/golisp/utils/enforce"
	"github.com/lucasew/golisp/vm"
)

var isFuncs = stdlib.NewRepository()

func GetRepo() stdlib.Repository {
	if !isFuncs.IsNil() {
		return isFuncs
	}
	for _, e := range register.Registry {
		Register(e)
	}
	return isFuncs
}

func Register(e data.LispEntity) {
	f := func() interface{} {
		return func(ctx context.Context, v ...data.LispValue) (data.LispValue, error) {
			err := enforce.Validate(
				enforce.Length(v, 1),
			)
			if err != nil {
				return types.Nil, err
			}
			if e.EntityIsFn(v[0]) {
				return types.T, nil
			}
			return types.Nil, nil
		}
	}
	isFuncs.Register("is", fmt.Sprintf("is-%s", e.EntityName()), f)
}

func ImportOnVM(vm vm.LispVM) {
	GetRepo().ImportOnVM(vm, "is")
}
