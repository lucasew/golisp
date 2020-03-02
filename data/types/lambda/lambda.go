package lambda

import (
	"github.com/lucasew/golisp/data"
	"github.com/lucasew/golisp/vm"
)

type lispLambda struct {
	vm     vm.LispVM
	ast    data.LispCons
	params []string
}

func NewLambda(vm vm.LispVM, ast data.LispCons, params []string) data.LispFunction {
	return lispLambda{
		vm,
		ast,
		params,
	}
}

func (f lispLambda) IsNil() bool {
	return false
}

func (f lispLambda) LispCall(i ...data.LispValue) (data.LispValue, error) {
	vm := f.vm.PushVM()
	for k := range f.params {
		vm.EnvSetLocal(f.params[k], i[k])
	}
	return vm.Eval(f.ast)
}

func (f lispLambda) Repr() string {
	return "<lambda function>"
}

func (f lispLambda) IsFunctionNative() bool {
	return false
}

func (lispLambda) LispTypeName() string {
	return "function"
}
