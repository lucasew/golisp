package convert

import (
    "errors"
    "reflect"
    "fmt"

    "github.com/lucasew/golisp/data/types"
    "github.com/lucasew/golisp/data"
    "github.com/lucasew/golisp/data/macro"
    "github.com/lucasew/golisp/vm"
)

var ErrCantHandleType = errors.New("cant handle type")

func NewLispList(v ...interface{}) (lv data.LispValue, err error) {
    lv = data.Nil
    ret := make([]data.LispValue, len(v))
    for k, v := range v {
        ret[k], err = NewLispValue(v)
        if err != nil {
            return
        }
    }
    lv = types.NewCons(ret...)
    return
}

func NewLispValue(v interface{}) (lv data.LispValue, err error) {
    t, ok := v.(data.LispValue)
    if ok {
        lv = t
        return
    }
    lv = data.Nil
    switch i:= v.(type) {
    case string:
        lv = types.NewConventionalString(i)
    case nil:
        lv = data.Nil
    case bool:
        if i {
            lv = data.T
        } else {
            lv = data.Nil
        }
    case byte:
        lv = types.NewByte(i)
    case float64:
        lv = types.NewFloatFromFloat64(i)
    case int64:
        lv = types.NewIntFromInt64(i)
    case int, int16, int32, uint16, uint32:
        lv = types.NewIntFromInt64(reflect.ValueOf(i).Int())
    case func(data.LispCons)(data.LispValue, error):
        lv = types.NewFunction(i)
    case func(vm.LispVM, data.LispCons) (data.LispValue, error):
        lv = macro.NewLispMacro(i)
    case data.IntoLispValue:
        lv = i.ToLispValue()
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
    err = fmt.Errorf("%w: %T", ErrCantHandleType, v)
    return
}
