package types

import (
	"fmt"
	"github.com/lucasew/golisp/data"
)

type Byte byte

func NewByte(b byte) Byte {
	return Byte(b)
}

func (Byte) IsNil() bool {
	return false
}

func (b Byte) Repr() string {
	return fmt.Sprintf("#c%d", b) // TODO: Improve
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

func (Byte) LispTypeName() string {
	return "byte"
}

func (b Byte) Sum(other Byte) Byte {
	return b + other
}

func (b Byte) Sub(other Byte) Byte {
	return b - other
}
