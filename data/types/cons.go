package types

import (
    "strings"
    "fmt"
    "github.com/lucasew/golisp/data"
)

type Cons []data.LispValue

func NewCons(v ...data.LispValue) data.LispCons {
    return Cons(v)
}

func (i Cons) Car() data.LispValue {
    if len(i) == 0 {
        return data.Nil
    }
    return i[0]
}

func (i Cons) Cdr() data.LispValue {
    if len(i) < 2 {
        return data.Nil
    }
    return i[1:len(i)]
}

func (i Cons) IsNil() bool {
    return len(i) == 0
}

func (i Cons) Repr() string {
    strs := make([]string, len(i))
    for k, v := range i {
        strs[k] = v.Repr()
    }
    return fmt.Sprintf(" (%s) ", strings.Join(strs, " "))
}

func (i Cons) Get(k int) data.LispValue {
    if k > len(i) {
        return data.Nil
    }
    return i[k]
}

func (i Cons) Len() int {
    return len(i)
}
