package vm_default

import (
	"errors"
	"github.com/lucasew/golisp/data"
	"github.com/lucasew/golisp/data/macro"
	"github.com/lucasew/golisp/data/types"
	common "github.com/lucasew/golisp/vm"
	"github.com/lucasew/golisp/vm/components/env"
	"github.com/lucasew/golisp/vm/default/stdlib"
)

type LispVM struct {
	env *env.LispEnv
}

func NewVM(e *env.LispEnv) common.LispVM {
	vm := LispVM{
		env: stdlib.NewDefaultEnv(e),
	}
	return &vm
}

func (vm *LispVM) PushVM() *LispVM {
	return &LispVM{
		env: env.NewLispEnv(vm.env),
	}
}

func (vm *LispVM) EnvGet(k string) data.LispValue {
	return vm.env.Get(k)
}

func (vm *LispVM) EnvSetLocal(k string, v data.LispValue) data.LispValue {
	return vm.env.SetLocal(k, v)
}

func (vm *LispVM) EnvSetGlobal(k string, v data.LispValue) data.LispValue {
	return vm.env.SetGlobal(k, v)
}

// Eval this function is where the magic starts
func (vm *LispVM) Eval(v data.LispValue) (data.LispValue, error) {
	switch in := v.(type) {
	case types.Cons:
		switch first := in.Car().(type) {
		case types.Symbol:
			f := vm.EnvGet(first.ToString())
			switch fn := f.(type) {
			case data.LispFunction:
				crude_params := []data.LispValue(in.Cdr().(types.Cons))
				params := make([]data.LispValue, len(crude_params))
				var err error = nil
				for k, v := range crude_params {
					params[k], err = vm.Eval(v)
					if err != nil {
						return types.Nil, err
					}
				}
				return fn.LispCall(types.NewCons(params...))
			case macro.LispMacro:
				l, ok := in.Cdr().(data.LispCons)
				if !ok {
					l = types.NewCons(v)
				}
				return fn.LispCallMacro(vm, l)
			default:
				return types.Nil, errors.New("cant call the variable")
			}
		case types.Cons:
			ret := types.Nil.(data.LispValue)
			var err error
			for _, stmt := range in {
				ret, err = vm.Eval(stmt.(data.LispValue))
				if err != nil {
					return types.Nil, err
				}
			}
			return ret, nil
		default:
			println(v.Repr())
			return types.Nil, errors.New("in code mode only commands are allowed")
		}
	case types.Symbol:
		return vm.EnvGet(in.ToString()), nil
	case types.String:
		return in, nil
	default:
		return in, nil
	}
}
