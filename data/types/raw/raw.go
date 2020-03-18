package raw

import (
	"context"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/lucasew/golisp/data"
	"github.com/lucasew/golisp/data/entity"
	"github.com/lucasew/golisp/data/entity/register"
	"github.com/lucasew/golisp/data/types"
	"github.com/lucasew/golisp/data/types/macro"
	"github.com/lucasew/golisp/data/types/maps"
	"github.com/lucasew/golisp/data/types/number"
	"github.com/lucasew/golisp/vm"
	"math/big"
	"reflect"
)

func init() {
	register.Register(new(LispWrapper).LispEntity())
}

func (LispWrapper) LispEntity() data.LispEntity {
	return entity.Entity{
		"wrapper", func(v data.LispValue) bool {
			_, ok := v.(LispWrapper)
			return ok
		},
	}
}

type LispWrapper struct {
	v interface{}
}

func NewLispWrapper(v interface{}) data.LispValue {
	t, ok := v.(data.LispValue)
	if ok {
		return t
	}
	u, ok := v.(data.IntoLispValue)
	if ok {
		return u.ToLispValue()
	}
	switch reflect.TypeOf(v).Kind() {
	case reflect.Slice, reflect.Array:
		return NewConsWrapper(v)
	case reflect.String:
		return types.NewString(v.(string))
	case reflect.Chan:
		return NewPortalFromNativeChan(v)
	case reflect.Bool:
		if v.(bool) {
			return types.T
		} else {
			return types.Nil
		}
	}
	switch v := v.(type) {
	case map[string]data.LispValue:
		return maps.NewMapFromMapString(v)
	case map[data.LispValue]data.LispValue:
		return maps.NewMapValue(v)
	case int, int16, int32, int64, int8, uint, uint16, uint32, uint8:
		return number.NewIntFromInt64(reflect.ValueOf(v).Int())
	case uint64:
		// TODO
		return types.Nil
	case float32:
		return number.NewFloatFromFloat64(float64(v))
	case float64:
		return number.NewFloatFromFloat64(v)
	case *big.Int:
		return number.NewIntFromBigInt(v)
	case *big.Float:
		return number.NewFloatFromBigFloat(v)
	case *big.Rat:
		return number.NewRationalFromBigRat(v)
	case big.Int, big.Float, big.Rat:
		return NewLispWrapper(&v)
	case func(context.Context, vm.LispVM, ...data.LispValue) (data.LispValue, error):
		return macro.NewLispMacro(v)
	case func(context.Context, ...data.LispValue) (data.LispValue, error):
		return types.NewFunction(v)
	}
	return LispWrapper{v}
}

func (lw LispWrapper) IsNil() bool {
	if lw.v == nil {
		return true
	}
	if lw.Len() == 0 {
		return true
	}
	return false
}

func (lw LispWrapper) Len() int {
	l, ok := lw.v.(interface{ Len() int })
	if ok {
		return l.Len()
	}
	v := reflect.ValueOf(lw.v)
	switch reflect.TypeOf(lw.v).Kind() {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		return v.Len()
	}
	return 1
}

func (lw LispWrapper) ToString() string {
	switch v := lw.v.(type) {
	case string:
		return v
	case data.LispString:
		return v.ToString()
	case interface{ String() string }:
		return v.String()
	}
	return fmt.Sprintf("%+v", lw.v)
}

func (lw LispWrapper) LispTypeName() string {
	return fmt.Sprintf("native.%T", lw.v)
}

func (lw LispWrapper) Repr() string {
	return fmt.Sprintf("<native %s >", spew.Sdump(lw.v))
}
