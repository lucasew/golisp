package vm

import (
    "testing"
    "github.com/lucasew/golisp/datatypes"
)

func TestEnvGetSet(t *testing.T) {
    e := NewLispEnv(nil)
    l, _ := datatypes.NewLispList(1, "2", 3.3)
    e.SetLocal("valor", l)
    if e.Get("valor").Repr() != l.Repr() {
        t.Fail()
    }
}

func TestHierarquicalEnvSet(t *testing.T) {
    a := NewLispEnv(nil)
    l, _ := datatypes.NewLispList(1, "2", 3.3)
    a.SetLocal("valor", l)
    b := NewLispEnv(a)
    if b.Get("valor").Repr() != l.Repr() {
        t.Fail()
    }
}
