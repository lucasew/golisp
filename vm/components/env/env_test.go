package env

import (
	"github.com/lucasew/golisp/data/types"
	"github.com/lucasew/golisp/data/types/number"
	"testing"
)

func TestEnvGetSet(t *testing.T) {
	e := NewLispEnv(nil)
    l := types.NewCons(
        number.NewIntFromInt(1),
        types.NewString("2"),
        number.NewFloatFromFloat64(3.3),
    )
	e.SetLocal("valor", l)
	if e.Get("valor").Repr() != l.Repr() {
		t.Fail()
	}
}

func TestHierarquicalEnvSet(t *testing.T) {
	a := NewLispEnv(nil)
    l := types.NewCons(
        number.NewIntFromInt(1),
        types.NewString("2"),
        number.NewFloatFromFloat64(3.3),
    )
	a.SetLocal("valor", l)
	b := NewLispEnv(a)
	if b.Get("valor").Repr() != l.Repr() {
		t.Fail()
	}
}
