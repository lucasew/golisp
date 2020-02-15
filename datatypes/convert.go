package datatypes

import (
    "errors"
    "reflect"
    "fmt"
)

func NewLispList(v ...interface{}) (lv LispValue, err error) {
    lv = Nil
    ret := make([]LispValue, len(v))
    for k, v := range v {
        ret[k], err = NewLispValue(v)
        if err != nil {
            return
        }
    }
    lv = NewCons(ret...)
    return
}

func NewLispValue(v interface{}) (lv LispValue, err error) {
    t, ok := v.(LispValue)
    if ok {
        lv = t
        return
    }
    lv = Nil
    switch i:= v.(type) {
    case string:
        lv = NewConventionalString(i)
    case nil:
        lv = Nil
    case bool:
        if i {
            lv = T
        } else {
            lv = Nil
        }
    case byte:
        lv = NewByte(i)
    case float64:
        lv = NewFloatFromFloat64(i)
    case int64:
        lv = NewIntFromInt64(i)
    case int, int16, int32, uint16, uint32:
        lv = NewIntFromInt64(reflect.ValueOf(i).Int())
    case func(LispValue)(LispValue, error):
        lv = NewFunction(i)
    default:
        goto checkvec
    }
    return
    checkvec:
    if reflect.TypeOf(v).Kind() == reflect.Slice {
        i := v
        l := reflect.ValueOf(i).Len()
        ret := make([]interface{}, l)
        for k := range ret {
            ret[k] = reflect.ValueOf(i).Index(k).Interface()
        }
        lv, err = NewLispList(ret...)
        return
    }
    err = errors.New(fmt.Sprintf("cant handle type %T", v))
    return
}
