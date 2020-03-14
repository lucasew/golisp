package params

import (
    "github.com/lucasew/golisp/data/types/test"
    "github.com/lucasew/golisp/data"
    "fmt"
)

var criterias = map[string]func(data.LispValue) bool {
    "atom": test.IsAtom,
    "carcdr": test.IsCarCdr,
    "cons": test.IsCons,
    "function": test.IsFunction,
    "iterator": test.IsIterator,
    "len": test.IsLen,
    "map": test.IsMap,
    "namespace": test.IsNamespace,
    "number": test.IsNumber,
    "portal": test.IsPortal,
    "string": test.IsString,
    "value": test.IsValue,
}

func RegisterCriteria(k string, v func(data.LispValue)bool) {
    criterias[k] = v
}

func CriteriaByName(name string) func(data.LispValue) error {
    return NewCriteriaForIsFuncion(criterias[name], name)
}

func CriteriaSameType(v data.LispValue) func(data.LispValue) error {
    return func (w data.LispValue) error {
        if !test.IsSameType(v, w) {
            return fmt.Errorf("not the raw same type: %s vs %s", v.LispTypeName(), w.LispTypeName())
        }
        return nil
    }
}
