package test

import (
	"github.com/lucasew/golisp/data"
	"github.com/lucasew/golisp/data/entity"
	"github.com/lucasew/golisp/data/entity/register"
	"testing"
)

func EntityAsTest(e entity.Entity) func(data.LispValue, *testing.T) {
	return func(v data.LispValue, t *testing.T) {
		t.Run(e.Name, func(t *testing.T) {
			if e.Isfn == nil {
				t.Errorf("entity %s have nil is function", e.Name)
			}
			if !e.Isfn(v) {
				t.Fail()
			}
		})
	}
}

func TestValues(v data.LispValue, t *testing.T, entities ...string) {
	for _, name := range entities {
		println(name)
		e, ok := register.Get(name)
		if !ok {
			t.Errorf("invalid entity tested")
			continue
		}
		EntityAsTest(e)(v, t)
	}
}
