package types

import (
	"github.com/lucasew/golisp/data"
	"github.com/lucasew/golisp/data/types/number"
	"testing"
)


func NumberTest(v data.LispValue) func(t *testing.T) {
	return func(t *testing.T) {
		if !number.IsNumber(v) {
			t.Fail()
		}
	}
}
