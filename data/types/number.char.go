package types

import (
    "fmt"
    "github.com/lucasew/golisp/data"
    "testing"
)

type Byte byte

func NewByte(b byte) Byte {
    return Byte(b)
}

func (Byte) IsNil() bool {
    return false
}

func (b Byte) Repr() string {
    return fmt.Sprintf("b%d", b) // TODO: Improve
}

func (b Byte) IsInfinite() bool {
    return false
}

func (b Byte) IsInt() bool {
    return true
}

func (b Byte) IsZero() bool {
    return b == 0
}

func IsByte(v data.LispValue) bool {
    _, ok := v.(Byte)
    return ok
}

func ByteTest(v data.LispValue) func(*testing.T) {
    return func(t *testing.T) {
        if !IsByte(v) {
            t.Fail()
        }
    }
}

func (Byte) LispTypeName() string {
    return "byte"
}
