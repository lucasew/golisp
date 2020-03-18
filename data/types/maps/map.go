package maps

import (
	"fmt"
	"github.com/lucasew/golisp/data"
	"github.com/lucasew/golisp/data/entity"
	"github.com/lucasew/golisp/data/entity/register"
	"github.com/lucasew/golisp/data/types"
	"strings"
)

func init() {
	register.Register(new(LispMap).LispEntity())
}

func (LispMap) LispEntity() data.LispEntity {
	return entity.Entity{
		"value_map", func(v data.LispValue) bool {
			_, ok := v.(LispMap)
			return ok
		},
	}
}

func NewMapValue(m map[data.LispValue]data.LispValue) LispMap {
	return LispMap(m)
}

type LispMap map[data.LispValue]data.LispValue

func NewMap() data.LispMap {
	return LispMap(map[data.LispValue]data.LispValue{})
}

func (m LispMap) Unwrap() map[data.LispValue]data.LispValue {
	return map[data.LispValue]data.LispValue(m)
}

func (m LispMap) Get(k data.LispValue) data.LispValue {
	r, ok := m[k]
	if !ok {
		return types.Nil
	}
	return r
}

func (m LispMap) Set(k data.LispValue, v data.LispValue) data.LispValue {
	m[k] = v
	return v
}

func (m LispMap) Keys() data.LispCons {
	ret := make([]data.LispValue, len(m))
	i := 0
	for k := range m {
		ret[i] = k
		i++
	}
	return types.Cons(ret)
}

func (m LispMap) Values() data.LispCons {
	ret := make([]data.LispValue, len(m))
	i := 0
	for _, v := range m {
		ret[i] = v
		i++
	}
	return types.Cons(ret)
}

func (m LispMap) Tuples() data.LispCons {
	ret := make([]data.LispValue, len(m))
	i := 0
	for k, v := range m {
		ret[i] = types.Cons([]data.LispValue{k, v})
		i++
	}
	return types.Cons(ret)
}

func (m LispMap) Len() int {
	return len(m)
}

func (m LispMap) LispTypeName() string {
	return "map"
}

func (m LispMap) Repr() string {
	ret := []string{"( new-map "}
	for k, v := range m {
		kf := ""
		switch ks := k.(type) {
		case data.LispString:
			kf = ks.ToString()
		default:
			kf = fmt.Sprintf(":invalid[%s]", ks.Repr())
		}
		ret = append(ret, fmt.Sprintf("%s '%s ", kf, v.Repr()))
	}
	ret = append(ret, ")")
	return strings.Join(ret, " ")
}

func (m LispMap) IsNil() bool {
	return len(m) == 0
}
