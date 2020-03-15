package repo

import (
	"fmt"
	"github.com/lucasew/golisp/data/entity"
	"github.com/lucasew/golisp/data/entity/lispfunc"
	"github.com/lucasew/golisp/data/entity/register"
	"github.com/lucasew/golisp/stdlib"
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

func Register(e entity.Entity) {
	f := func() interface{} {
		return lispfunc.EntityAsLispFunc(e)
	}
	isFuncs.Register("is", fmt.Sprintf("is-%s", e.Name), f)
}

func ImportOnVM(vm vm.LispVM) {
	GetRepo().ImportOnVM(vm, "is")
}
