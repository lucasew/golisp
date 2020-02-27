package lambda

import (
	"github.com/lucasew/golisp/data"
	"github.com/lucasew/golisp/data/types/iterator"
	"github.com/lucasew/golisp/vm"
)

type lispLambda struct {
    vm vm.LispVM
    ast data.LispCons
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

func (f lispLambda) LispCall(i data.LispCons) (data.LispValue, error) {
    vm := f.vm.PushVM()
    params := iterator.NewConsIterator(i)
    for _, v := range f.params {
        vm.EnvSetLocal(v, params.Next())
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

