package stdlib

import (
    "github.com/davecgh/go-spew/spew"
    "github.com/lucasew/golisp/data"
    "github.com/lucasew/golisp/data/types"
    "github.com/lucasew/golisp/vm"
)

func init() {
    register("env-dump", EnvDump)
}

func EnvDump(env vm.LispVM, v data.LispCons) (data.LispValue, error) {
    return types.NewConventionalString(spew.Sdump(env)), nil
}
