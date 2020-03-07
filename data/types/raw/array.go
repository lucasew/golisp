package raw

import (
	"github.com/lucasew/golisp/data"
	"github.com/lucasew/golisp/data/types"
	"reflect"
)

type LispConsWrapper struct {
	v interface{}
}

func NewConsWrapper(v interface{}) data.LispValue {
	if reflect.TypeOf(v).Kind() != reflect.Slice {
		if reflect.TypeOf(v).Kind() != reflect.Array {
			return nil
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
	return "< todo lisp cons wrapper >"
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
		return NewLispWrapper(reflect.ValueOf(c.v).Index(k))
	}
	return types.Nil
}

func init() {
	var v interface{} = NewLispWrapper([]int{1})
	_ = v.(data.LispCons)
}
