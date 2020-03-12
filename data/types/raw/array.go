package raw

import (
	"fmt"
	"github.com/lucasew/golisp/data"
	"github.com/lucasew/golisp/data/types"
	"reflect"
	"strings"
)

type LispConsWrapper struct {
	v interface{}
}

func NewConsWrapper(v interface{}) data.LispValue {
	if reflect.TypeOf(v).Kind() != reflect.Slice {
		if reflect.TypeOf(v).Kind() != reflect.Array {
			return types.Nil
		}
	}
	return LispConsWrapper{
		v: v,
	}
}

func (c LispConsWrapper) Len() int {
	return reflect.ValueOf(c.v).Len()
}

func (c LispConsWrapper) IsNil() bool {
	return c.Len() == 0
}

func (c LispConsWrapper) LispTypeName() string {
	return "cons"
}

func (c LispConsWrapper) Repr() string {
	ret := make([]string, c.Len())
	for k := range ret {
		ret[k] = c.Get(k).Repr()
	}
	return fmt.Sprintf("( %s )", strings.Join(ret, " "))
}

func (c LispConsWrapper) Car() data.LispValue {
	if c.Len() > 0 {
		return NewLispWrapper(reflect.ValueOf(c.v).Index(0).Interface())
	}
	return types.Nil
}

func (c LispConsWrapper) Cdr() data.LispCarCdr {
	if c.Len() > 0 {
		ret, ok := NewLispWrapper(reflect.ValueOf(c.v).Slice(1, c.Len()).Interface()).(data.LispCarCdr)
		if ok {
			return ret
		} else {
			return types.Nil
		}
	}
	return types.Nil
}

func (c LispConsWrapper) Get(k int) data.LispValue {
	if c.Len() > k {
		return NewLispWrapper(reflect.ValueOf(c.v).Index(k).Interface())
	}
	return types.Nil
}

func init() {
	var v interface{} = NewLispWrapper([]int{1})
	_ = v.(data.LispCons)
}
