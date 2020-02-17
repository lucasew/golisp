package vm_default

import (
    common "github.com/lucasew/golisp/vm"
    "github.com/lucasew/golisp/vm/components/env"
    "github.com/lucasew/golisp/data"
    "github.com/lucasew/golisp/data/types"
    "github.com/lucasew/golisp/data/macro"
    "errors"
    // "github.com/lucasew/golisp/parser"
    // "fmt"
)

type LispVM struct {
    env *env.LispEnv
}


func NewVM(e *env.LispEnv) common.LispVM {
    vm := LispVM{
        env: NewDefaultEnv(e),
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
    // _, literal := v.Car().(parser)
    // if literal {
    //     return v, nil
    // }
    switch in := v.(type) {
    case types.Cons:
        if in.Len() == 1 {
            return vm.Eval(v.Car())
        }
        switch first := in.Car().(type) {
        case types.Symbol:
            f := vm.EnvGet(first.ToString())
            switch fn := f.(type) {
            case data.LispFunction:
                crude_params := []data.LispValue(in.Cdr().(types.Cons))
                params := make([]data.LispValue, len(crude_params))
                var err error = nil
                for k, v := range crude_params {
                    params[k], err = vm.PushVM().Eval(v)
                    if err != nil {
                        return data.Nil, err
                    }
                }
                return fn.LispCall(types.NewCons(params...))
            case macro.LispMacro:
                return fn.LispCallMacro(vm.PushVM(), v.Cdr())
            default:
                return data.Nil, errors.New("cant call the variable")
            }
        default:
            println(v.Repr())
            return data.Nil, errors.New("in code mode only commands are allowed")
        }
    case types.Symbol:
        return vm.EnvGet(in.ToString()), nil
    case types.ConventionalString:
        return in, nil
    default:
        return in, nil
    }
}
