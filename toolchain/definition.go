package toolchain

import (
	"context"
	"github.com/lucasew/golisp/data"
	"github.com/lucasew/golisp/data/types"
	"github.com/lucasew/golisp/vm"
)

type Toolchain struct {
	vm.LispVM
	parsefn func(context.Context, string) (data.LispValue, error)
}

func NewToolchain(VM vm.LispVM, parsefn func(context.Context, string) (data.LispValue, error)) Toolchain {
	return Toolchain{
		VM, parsefn,
	}
}

func (t Toolchain) ParseString(ctx context.Context, stmt string) (data.LispValue, error) {
	return t.parsefn(ctx, stmt)
}

func (t Toolchain) ParseBytes(ctx context.Context, stmt []byte) (data.LispValue, error) {
	return t.parsefn(ctx, string(stmt))
}

func (t *Toolchain) EvalString(ctx context.Context, stmt string) (data.LispValue, error) {
	pstmt, err := t.parsefn(ctx, stmt)
	if err != nil {
		return types.Nil, err
	}
	return t.Eval(ctx, pstmt)
}

func (t *Toolchain) EvalBytes(ctx context.Context, stmt []byte) (data.LispValue, error) {
	return t.EvalString(ctx, string(stmt))
}
