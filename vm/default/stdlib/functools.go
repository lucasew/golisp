package stdlib

import (
    "github.com/lucasew/golisp/data"
    "github.com/lucasew/golisp/data/types"
    "fmt"
)

func init() {
    register("map", Map)
    register("reduce", Reduce)
}

func Map(v data.LispCons) (data.LispValue, error) {
    fn, ok := v.Car().(data.LispFunction)
    if !ok {
        return data.Nil, fmt.Errorf("first argument must be a function, got %T", v.Car())
    }
    lst, ok := v.Cdr().Car().(data.LispCons)
    if !ok {
        return data.Nil, fmt.Errorf("second argument must be a cons, got %T", v.Cdr().Car())
    }
    var err error
    ret := make([]data.LispValue, lst.Len())
    for i := 0; i < lst.Len(); i++ {
        v := types.NewCons(lst.Get(i))
        ret[i], err = fn.LispCall(v)
        if err != nil {
            return data.Nil, err
        }
    }
    return types.NewCons(ret...), nil
}

func Reduce(v data.LispCons) (data.LispValue, error) {
    fn, ok := v.Car().(data.LispFunction)
    if !ok {
        return data.Nil, fmt.Errorf("first argument must be a function, got %T", v.Car())
    }
    lst, ok := v.Cdr().Car().(data.LispCarCdr)
    if !ok {
        return data.Nil, fmt.Errorf("second argument must be a cons, got %T", v.Cdr().Car())
    }
    ret := lst.Car()
    next:
    if lst.Cdr().IsNil() {
        return ret, nil
    }
    lst = lst.Cdr()
    ret, err := fn.LispCall(types.NewCons(ret, lst.Car()))
    if err != nil {
        return data.Nil, err
    }
    goto next
}
