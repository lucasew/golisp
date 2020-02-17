package datatypes

import (
    "strings"
    "fmt"
)

type Map map[LispValue]LispValue

func NewMapWith(things map[interface{}]interface{}) (lv LispValue, err error) {
    lv = Nil
    err = nil
    var mapret LispMap = nil
    ret := map[LispValue]LispValue{}
    var ck LispValue
    var cv LispValue
    for k, v := range things {
        ck, err = NewLispValue(k)
        if err != nil {
            return
        }
        cv, err = NewLispValue(v)
        if err != nil {
            return
        }
        ret[ck] = cv
    }
    mapret = Map(ret)
    lv = mapret
    return
}

func (Map) Car() LispValue {
    return Nil
}

func (Map) Cdr() LispValue {
    return Nil
}

func (m Map) Get(k LispValue) LispValue {
    v, ok := m[k]
    if !ok {
        return Nil
    }
    return v
}

func (m Map) Set(k LispValue, v LispValue) LispValue {
    m[k] = v
    return v
}

func (m Map) Keys() LispValue {
    ret := make([]LispValue, len(m))
    i := 0
    for k := range(m) {
        ret[i] = k
        i++
    }
    return NewCons(ret...)
}

func (m Map) Values() LispValue {
    ret := make([]LispValue, len(m))
    i := 0
    for _, v := range(m) {
        ret[i] = v
        i++
    }
    return NewCons(ret...)
}

func (m Map) Tuples() LispValue {
    ret := make([]LispValue, len(m))
    i := 0
    for k, v := range(m) {
        ret[i] = NewCons(k, v)
        i++
    }
    return NewCons(ret...)
}

func (m Map) IsNil() bool {
    return len(m) == 0
}

func (m Map) Len() int {
    return len(m)
}

func (m Map) Repr() string {
    ret := make([]string, len(m))
    i := 0
    for k, v := range m {
        ret[i] = fmt.Sprintf("( %s : %s )", k.Repr(), v.Repr())
        i++
    }
    return fmt.Sprintf("( %s )", strings.Join(ret, " "))
}
