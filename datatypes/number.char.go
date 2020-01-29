package datatypes

import (
    "fmt"
)

type Byte byte

func NewByte(b byte) Byte {
    return Byte(b)
}

func (Byte) Car() LispValue {
    return Nil
}

func (Byte) Cdr() LispValue {
    return Nil
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
