package vm_default

import (
	"context"
	"errors"
	"github.com/lucasew/golisp/data"
	"github.com/lucasew/golisp/data/entity/repo"
	"github.com/lucasew/golisp/data/types"
	"github.com/lucasew/golisp/data/types/macro"
	"github.com/lucasew/golisp/data/types/raw"
	"github.com/lucasew/golisp/stdlib/default"
	common "github.com/lucasew/golisp/vm"
	"github.com/lucasew/golisp/vm/components/env"
)

type LispVM struct {
	env      *env.LispEnv
	imported *env.LispEnv
}

func NewVM(e *env.LispEnv) common.LispVM {
	imported := env.NewLispEnv(env.NewLispEnv(nil))
	vm := LispVM{
		env:      env.NewLispEnv(imported),
		imported: imported,
	}
	stdlib.ELEMENTS.ImportOnVM(&vm, "default")
	repo.ImportOnVM(&vm)
	return &vm
}

func (vm *LispVM) PushVM() common.LispVM {
	return &LispVM{
		env: env.NewLispEnv(vm.env),
	}
}

func (vm *LispVM) Import(m map[string]interface{}) {
	for k, v := range m {
		v := raw.NewLispWrapper(v)
		vm.imported.SetLocal(k, v)
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
func (vm *LispVM) Eval(ctx context.Context, v data.LispValue) (data.LispValue, error) {
	select {
	case _ = <-ctx.Done():
		return types.Nil, data.ErrContextCancelled
	default:
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
						params[k], err = vm.Eval(ctx, v)
						if err != nil {
							return types.Nil, err
						}
					}
					return fn.LispCall(ctx, params...)
				case macro.LispMacro:
					l, ok := in.Cdr().(data.LispCons)
					if !ok {
						l = types.NewCons(v)
					}
					switch c := l.(type) {
					case types.Cons:
						return fn.LispCallMacro(ctx, vm, c...)
					default:
						return types.Nil, errors.New("evaluation of not []LispValue cons not supported yet")
					}
				default:
					return types.Nil, errors.New("cant call the variable")
				}
			case types.Cons:
				ret := types.Nil.(data.LispValue)
				var err error
				for _, stmt := range in {
					ret, err = vm.Eval(ctx, stmt.(data.LispValue))
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
}
