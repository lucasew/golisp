package macro

import (
	"context"
	"github.com/lucasew/golisp/data"
	"github.com/lucasew/golisp/data/entity/register"
	"github.com/lucasew/golisp/vm"
)

func init() {
	register.Register("macro", func(v data.LispValue) bool {
		_, ok := v.(LispMacro)
		return ok
	})
}

type LispMacro func(context.Context, vm.LispVM, ...data.LispValue) (data.LispValue, error)

func NewLispMacro(f func(context.Context, vm.LispVM, ...data.LispValue) (data.LispValue, error)) LispMacro {
	return f
}

func (LispMacro) IsNil() bool {
	return false
}

func (LispMacro) Repr() string {
	return "<native macro>"
}

func (LispMacro) LispTypeName() string {
	return "macro"
}

func (m LispMacro) LispCallMacro(ctx context.Context, v vm.LispVM, val ...data.LispValue) (data.LispValue, error) {
	return m(ctx, v, val...)
}
