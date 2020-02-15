package vm_default

import (
    common "github.com/lucasew/golisp/vm"
    "github.com/lucasew/golisp/datatypes"
    // "github.com/lucasew/golisp/parser"
    "fmt"
)

type LispVM struct {
    env *common.LispEnv
}

func NewVM(env *common.LispEnv) common.LispVM {
    env = common.NewLispEnv(env)
    vm := LispVM{
        env: NewDefaultEnv(env),
    }
    return &vm
}

func (vm *LispVM) EnvGet(k string) datatypes.LispValue {
    return vm.env.Get(k)
}

func (vm *LispVM) EnvSet(k string, v datatypes.LispValue) datatypes.LispValue {
    return vm.env.SetLocal(k, v)
}

// Eval this function is where the magic starts
func (vm *LispVM) Eval(v datatypes.LispValue) (datatypes.LispValue, error) {
    // _, literal := v.Car().(parser)
    // if literal {
    //     return v, nil
    // }

    cmd, ok := v.Car().(datatypes.Symbol)
    if ok {
        f := vm.EnvGet(cmd.ToString())
        fn, ok := f.(datatypes.LispFunction)
        if ok {
            return fn.LispCall(v.Cdr())
        } else {
            return datatypes.Nil, fmt.Errorf("not a function: %s (%s)", cmd.ToString(), f.Repr())
        }
    } else {
        cons, is_cons := v.(datatypes.Cons)
        if is_cons {
            ret := make([]datatypes.LispValue, len(cons))
            for k, v := range cons {
                r, err := vm.Eval(v)
                if err != nil {
                    return r, err
                }
                ret[k] = r
            }
            return datatypes.NewCons(ret...), nil
        } else {
            return v, nil
        }
    }
}
