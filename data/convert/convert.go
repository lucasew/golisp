package convert

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/lucasew/golisp/data"
	"github.com/lucasew/golisp/data/macro"
	"github.com/lucasew/golisp/data/types"
	"github.com/lucasew/golisp/data/types/number"
	"github.com/lucasew/golisp/vm"
)

var ErrCantHandleType = errors.New("cant handle type")

func NewLispList(v ...interface{}) (lv data.LispValue, err error) {
	lv = types.Nil
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
	lv = types.Nil
	switch i := v.(type) {
	case string:
		lv = types.NewString(i)
	case nil:
		lv = types.Nil
	case bool:
		if i {
			lv = types.T
		} else {
			lv = types.Nil
		}
	case byte:
		lv = number.NewByte(i)
	case float64:
		lv = number.NewFloatFromFloat64(i)
	case int64:
		lv = number.NewIntFromInt64(i)
	case int, int16, int32, uint16, uint32:
		lv = number.NewIntFromInt64(reflect.ValueOf(i).Int())
	case func(...data.LispValue) (data.LispValue, error):
		lv = types.NewFunction(i)
	case func(vm.LispVM, ...data.LispValue) (data.LispValue, error):
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
