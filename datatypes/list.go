package datatypes

import (
    "strings"
    "fmt"
)

type Cons []LispValue

func NewCons(v ...LispValue) LispCons {
    return Cons(v)
}

func (i Cons) Car() LispValue {
    if len(i) == 0 {
        return Nil
    }
    return i[0]
}

func (i Cons) Cdr() LispValue {
    if len(i) < 2 {
        return Nil
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

func (i Cons) Get(k int) LispValue {
    if k > len(i) {
        return Nil
    }
    return i[k]
}

func (i Cons) Len() int {
    return len(i)
}
