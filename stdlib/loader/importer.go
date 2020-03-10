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
