package mutexed

import (
	"context"
	"github.com/lucasew/golisp/data"
	vmdefs "github.com/lucasew/golisp/vm"
	"sync"
)

type MutexedVM struct {
	sync.Mutex
	vm vmdefs.LispVM
}

func NewVMMutexed(vm vmdefs.LispVM) vmdefs.LispVM {
	return &MutexedVM{
		vm: vm,
	}
}

func (e *MutexedVM) EnvGet(k string) data.LispValue {
	e.Lock()
	defer e.Unlock()
	return e.vm.EnvGet(k)
}

func (e *MutexedVM) EnvSetGlobal(k string, v data.LispValue) data.LispValue {
	e.Lock()
	defer e.Unlock()
	return e.vm.EnvSetGlobal(k, v)
}

func (e *MutexedVM) EnvSetLocal(k string, v data.LispValue) data.LispValue {
	e.Lock()
	defer e.Unlock()
	return e.vm.EnvSetLocal(k, v)
}

func (e *MutexedVM) Eval(ctx context.Context, v data.LispValue) (data.LispValue, error) {
	e.Lock()
	defer e.Unlock()
	return e.vm.Eval(ctx, v)
}

func (e *MutexedVM) Import(v map[string]interface{}) {
	e.Lock()
	defer e.Unlock()
	e.vm.Import(v)
}

func (e *MutexedVM) PushVM() vmdefs.LispVM {
	e.Lock()
	defer e.Unlock()
	return e.vm.PushVM()
}
