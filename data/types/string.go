package types

import (
    "fmt"
    "github.com/lucasew/golisp/data"
    "testing"
)

type String string

func NewString(s string) data.LispString {
    return String(s)
}

func (c String) IsNil() bool {
    return len(c) == 0
}

func (c String) Car() data.LispValue {
    if len(c) == 0 {
        return Nil
    }
    s := string(c)[0]
    return NewByte(s)
}

func (c String) Cdr() data.LispCarCdr {
    if len(c) < 2 {
        return Nil
    }
    s := string(c)[1:len(c)] // TODO: Testar se não teremos acessos errados aqui
    return NewString(s)
}

func (c String) Repr() string {
    return fmt.Sprintf("\"%s\"", c)
}

func (c String) ToString() string {
    return string(c)
}

func (String) LispTypeName() string {
    return "string"
}

func IsString(v data.LispValue) bool {
    _, ok := v.(data.LispString)
    return ok
}

func StringTest(v data.LispValue) func(*testing.T) {
    return func(t *testing.T) {
        if !IsString(v) {
            t.Fail()
        }
    }
}
