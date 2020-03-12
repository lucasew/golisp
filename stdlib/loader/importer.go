package loader

import (
    "github.com/lucasew/golisp/data/types/raw"
    "github.com/lucasew/golisp/data/types"
    "github.com/lucasew/golisp/utils/enforce"
    "github.com/lucasew/golisp/data"
    "github.com/lucasew/golisp/vm"
    "errors"
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

func (i Importer) IntoLispValue() data.LispValue {
    return raw.NewLispWrapper(
        func(vm vm.LispVM, v ...data.LispValue) (data.LispValue, error) {
            err := enforce.Validate(
                enforce.Length(v, 1),
                enforce.String(v, 1),
            )
            if err != nil {
                return types.Nil, err
            }
            module := v[0].(data.LispString).ToString()
            imported := i.ImportOnVM(vm, module)
            if len(imported) == 0 {
                return types.Nil, errors.New("module not found")
            }
            return raw.NewConsWrapper(imported), nil
        },
    )
}

func (i Importer) ImportOnVM(vm vm.LispVM, module string) []string {
    imported := []string{}
    for _, repo := range i.repositories {
        module, ok := repo.packages[module]
        if !ok {
            continue
        }
        to_import := map[string]interface{}{}
        for k, v := range module {
            imported = append(imported, k)
            to_import[k] = v()
        }
        vm.Import(to_import)
    }
    return imported
}

func (i *Importer) Keys() data.LispCons {
    ret := []string{}
    for _, repo := range i.repositories {
        for k, _ := range repo.packages {
            ret = append(ret, k)
        }
    }
    return raw.NewConsWrapper(ret).(data.LispCons)
}

func (i *Importer) Values() data.LispCons {
    ret := []data.LispValue{}
    for _, repo := range i.repositories {
        for _, v := range repo.packages {
            lv := raw.NewLispWrapper(v)
            ret = append(ret, lv)
        }
    }
    return raw.NewConsWrapper(ret).(data.LispCons)
}

func (i *Importer) Tuples() data.LispCons {
    ret := []data.LispValue{}
    for _, repo := range i.repositories {
        for k, v := range repo.packages {
            lv := raw.NewLispWrapper(v)
            ret = append(ret, types.NewCons(types.String(k), lv))
        }
    }
    return raw.NewConsWrapper(ret).(data.LispCons)
}

func (i *Importer) Get(lv data.LispValue) data.LispValue {
    key, ok := lv.(data.LispString)
    if !ok {
        return types.Nil
    }
    for _, repo := range i.repositories {
        for k, v := range repo.packages {
            if key.ToString() == k {
                return raw.NewLispWrapper(v)
            }
        }
    }
    return types.Nil
}

func (i *Importer) IsNil() bool {
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

func (i *Importer) Len() int {
    acc := 0
    for _, repo := range i.repositories {
        acc += len(repo.packages)
    }
    return acc
}

func (i *Importer) LispTypeName() string {
    return "importer"
}

func (i *Importer) Repr() string {
    return "< reporter >"
}

func (i *Importer) Set(k data.LispValue, v data.LispValue) data.LispValue {
    return types.Nil
}

func init() {
    var _ data.LispMap = NewImporter()
}
