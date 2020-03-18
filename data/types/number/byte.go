package number

import (
	"fmt"
	"github.com/lucasew/golisp/data"
	"github.com/lucasew/golisp/data/entity"
	"github.com/lucasew/golisp/data/entity/register"
)

func init() {
	register.Register(new(Byte).LispEntity())
}

func (Byte) LispEntity() data.LispEntity {
	return entity.Entity{
		"byte", func(v data.LispValue) bool {
			_, ok := v.(Byte)
			return ok
		},
	}
}

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

func (Byte) LispTypeName() string {
	return "byte"
}

func (b Byte) Sum(other Byte) Byte {
	return b + other
}

func (b Byte) Sub(other Byte) Byte {
	return b - other
}

func (b Byte) Cmp(other Byte) int {
	return int(b) - int(other)
}
