package types

import (
	"github.com/lucasew/golisp/data/entity/test"
	"testing"
)

func TestAtomType(t *testing.T) {
	v := NewAtom("")
	test.TestValues(v, t, "atom")
}

func TestAtomDecodeString(t *testing.T) {
	v := NewAtom("teste").(Atom)
	var str string
	if !v.LispEntity().AssignTo(v, &str) {
		t.Fail()
	}
	if string(str) != string(v) {
		t.Fail()
	}
}
