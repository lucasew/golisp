package types

import (
	"github.com/lucasew/golisp/data"
	"testing"
)

func NewTestHelper(f func(data.LispValue) bool) func(v data.LispValue) func(*testing.T) {
	return func(v data.LispValue) func(t *testing.T) {
		return func(t *testing.T) {
			if !f(v) {
				t.Fail()
			}
		}
	}
}
