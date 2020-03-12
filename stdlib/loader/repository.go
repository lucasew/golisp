package loader

import (
	"github.com/lucasew/golisp/data"
	"github.com/lucasew/golisp/vm"
)

type Repository struct {
	packages map[string]map[string]func() interface{}
}

func NewRepository() Repository {
	return Repository{
		map[string]map[string]func() interface{}{},
	}
}

func (r *Repository) Register(module string, name string, value func() interface{}) {
	if r.packages[module] == nil {
		r.packages[module] = map[string]func() interface{}{}
	}
	r.packages[module][name] = value
}

func (r Repository) ImportOnVM(vm vm.LispVM, module string) []string {
	return NewImporter(r).ImportOnVM(vm, module)
}

func (r Repository) IntoLispValue() data.LispValue {
	return NewImporter(r).IntoLispValue()
}

func init() {
	var _ data.LispMap = NewRepository()
}

func (r Repository) Get(k data.LispValue) data.LispValue {
	return NewImporter(r).Get(k)
}

func (r Repository) IsNil() bool {
	return NewImporter(r).IsNil()
}

func (r Repository) Keys() data.LispCons {
	return NewImporter(r).Keys()
}

func (r Repository) Len() int {
	return NewImporter(r).Len()
}

func (r Repository) LispTypeName() string {
	return "repository"
}

func (r Repository) Repr() string {
	return NewImporter(r).Repr()
}

func (r Repository) Set(k data.LispValue, v data.LispValue) data.LispValue {
	return NewImporter(r).Set(k, v)
}

func (r Repository) Tuples() data.LispCons {
	return NewImporter(r).Tuples()
}

func (r Repository) Values() data.LispCons {
	return NewImporter(r).Values()
}
