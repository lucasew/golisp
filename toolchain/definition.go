package toolchain

import (
	"github.com/lucasew/golisp/data"
	"github.com/lucasew/golisp/data/types"
	"github.com/lucasew/golisp/vm"
)

type Toolchain struct {
	vm.LispVM
	parsefn func(string) (data.LispValue, error)
}

func NewToolchain(VM vm.LispVM, parsefn func(string) (data.LispValue, error)) Toolchain {
	return Toolchain{
		VM, parsefn,
	}
}

func (t Toolchain) ParseString(stmt string) (data.LispValue, error) {
	return t.parsefn(stmt)
}

func (t Toolchain) ParseBytes(stmt []byte) (data.LispValue, error) {
	return t.parsefn(string(stmt))
}

func (t *Toolchain) EvalString(stmt string) (data.LispValue, error) {
	pstmt, err := t.parsefn(stmt)
	if err != nil {
		return types.Nil, err
	}
	return t.Eval(pstmt)
}

func (t *Toolchain) EvalBytes(stmt []byte) (data.LispValue, error) {
	return t.EvalString(string(stmt))
}
