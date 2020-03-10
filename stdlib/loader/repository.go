package loader

import (
    "github.com/lucasew/golisp/data"
    "github.com/lucasew/golisp/vm"
)

type Repository struct {
    packages map[string]map[string]func()interface{}
}

func NewRepository() Repository {
    return Repository{
        map[string]map[string]func()interface{}{},
    }
}

func (r *Repository) Register(module string, name string, value func()interface{}) {
    if r.packages[module] == nil {
        r.packages[module] = map[string]func()interface{}{}
    }
    r.packages[module][name] = value
}

func (r Repository) ImportOnVM(vm vm.LispVM, module string) []string {
    return NewImporter(r).ImportOnVM(vm, module)
}

func (r Repository) IntoLispValue() data.LispValue {
    return NewImporter(r).IntoLispValue()
}
