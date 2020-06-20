package stdlib

import (
	"github.com/lucasew/golisp/data"
	"github.com/lucasew/golisp/data/types"
	"github.com/lucasew/golisp/data/types/raw"
	"github.com/lucasew/golisp/vm"
)

type Importer struct {
	repositories []Repository
}

func NewImporter(r ...Repository) *Importer {
	return &Importer{r}
}

func (i *Importer) RegisterRepository(r Repository) {
	i.repositories = append(i.repositories, r)
}

func (i Importer) IsFunctionNative() bool {
	return true
}

func (i Importer) ImportOnVM(vm vm.LispVM, module string) []string {
	keys := []string{}
	for _, repo := range i.repositories {
		ikeys := repo.ImportOnVM(vm, module)
		for _, k := range ikeys {
			keys = append(keys, k)
		}
	}
	return keys
}

func (i Importer) Keys() data.LispCons {
	ret := []string{}
	for _, repo := range i.repositories {
		for _, pkgs := range repo.packages {
			for k := range pkgs {
				ret = append(ret, k)
			}
		}
	}
	return raw.NewConsWrapper(ret).(data.LispCons)
}

func (i Importer) Values() data.LispCons {
	ret := []data.LispValue{}
	for _, repo := range i.repositories {
		for _, pkgs := range repo.packages {
			for _, v := range pkgs {
				lv := raw.NewLispWrapper(v())
				ret = append(ret, lv)
			}
		}
	}
	return raw.NewConsWrapper(ret).(data.LispCons)
}

func (i Importer) Tuples() data.LispCons {
	ret := []data.LispValue{}
	for _, repo := range i.repositories {
		for _, pkgs := range repo.packages {
			for k, v := range pkgs {
				lv := raw.NewLispWrapper(v())
				ret = append(ret, types.NewCons(types.String(k), lv))
			}
		}
	}
	return raw.NewConsWrapper(ret).(data.LispCons)
}

func (i Importer) Get(lv data.LispValue) data.LispValue {
	key, ok := lv.(data.LispString)
	if !ok {
		return types.Nil
	}
	for _, repo := range i.repositories {
		for _, pkg := range repo.packages {
			for k, v := range pkg {
				if key.ToString() == k {
					return raw.NewLispWrapper(v())
				}
			}
		}
	}
	return types.Nil
}

func (i Importer) IsNil() bool {
	if len(i.repositories) == 0 {
		return true
	}
	for _, repo := range i.repositories {
		if len(repo.packages) > 0 {
			return false
		}
	}
	return true
}

func (i Importer) Len() int {
	acc := 0
	for _, repo := range i.repositories {
		for _, pkgs := range repo.packages {
			acc += len(pkgs)
		}
	}
	return acc
}

func (i Importer) LispTypeName() string {
	return "importer"
}

func (i Importer) Repr() string {
	return "< reporter >"
}

func (i Importer) Set(k data.LispValue, v data.LispValue) data.LispValue {
	return types.Nil
}

func init() {
	var _ data.LispMap = NewImporter()
	var _ data.LispValue = NewImporter()
}
