package vm_default

import (
    common "github.com/lucasew/golisp/vm"
    "github.com/lucasew/golisp/vm/components/env"
    "github.com/lucasew/golisp/data"
    "github.com/lucasew/golisp/data/types"
    // "github.com/lucasew/golisp/parser"
    "fmt"
)

type LispVM struct {
    env *env.LispEnv
}

func NewVM(e *env.LispEnv) common.LispVM {
    e = env.NewLispEnv(e)
    vm := LispVM{
        env: NewDefaultEnv(e),
    }
    return &vm
}

func (vm *LispVM) EnvGet(k string) data.LispValue {
    return vm.env.Get(k)
}

func (vm *LispVM) EnvSet(k string, v data.LispValue) data.LispValue {
    return vm.env.SetLocal(k, v)
}

// Eval this function is where the magic starts
func (vm *LispVM) Eval(v data.LispValue) (data.LispValue, error) {
    // _, literal := v.Car().(parser)
    // if literal {
    //     return v, nil
    // }

    cmd, ok := v.Car().(types.Symbol)
    if ok {
        f := vm.EnvGet(cmd.ToString())
        fn, ok := f.(data.LispFunction)
        if ok {
            return fn.LispCall(v.Cdr())
        } else {
            return data.Nil, fmt.Errorf("not a function: %s (%s)", cmd.ToString(), f.Repr())
        }
    } else {
        cons, is_cons := v.(types.Cons)
        if is_cons {
            ret := make([]data.LispValue, len(cons))
            for k, v := range cons {
                r, err := vm.Eval(v)
                if err != nil {
                    return r, err
                }
                ret[k] = r
            }
            return types.NewCons(ret...), nil
        } else {
            return v, nil
        }
    }
}
