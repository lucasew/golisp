package stdlib

import (
	"github.com/lucasew/golisp/data"
	"github.com/lucasew/golisp/data/types"
	"github.com/lucasew/golisp/data/types/macro"
	"github.com/lucasew/golisp/data/types/raw"
	"github.com/lucasew/golisp/utils/enforce"
	"github.com/lucasew/golisp/vm"
)

type Importable interface {
	ImportOnVM(vm.LispVM, string) []string
}

func AsMacro(i Importable) macro.LispMacro {
	return macro.LispMacro(
		func(vm vm.LispVM, v ...data.LispValue) (data.LispValue, error) {
			err := enforce.Validate(
				enforce.Length(v, 1),
				enforce.Entity("lisp_string", v, 1),
			)
			if err != nil {
				return types.Nil, err
			}
			pkg := v[0].(data.LispString).ToString()
			return raw.NewLispWrapper(i.ImportOnVM(vm, pkg)), nil
		},
	)
}
