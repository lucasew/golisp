package types

import (
	"github.com/lucasew/golisp/data/types/number"
	"github.com/lucasew/golisp/data/types/test"
	"testing"
)

func TestStringCar(t *testing.T) {
	s := NewString("hello")
	c := s.Car()
	if c != number.NewByte('h') {
		t.Errorf("expected 'h' got '%c'", c)
	}
}

func TestStringCdr(t *testing.T) {
	s := NewString("hello")
	c := s.Cdr()
	if c != NewString("ello") {
		t.Errorf("expected 'hello' got '%s'", c)
	}
	c = c.Cdr()
	if c != NewString("llo") {
		t.Errorf("expected 'hello' got '%s'", c)
	}
}

func TestStringCdrSmall(t *testing.T) {
	// let x be a string that (< (len x) 2)
	// -- (assert-true (eq (cdr x) nil))
	s := NewString("a")
	c := s.Cdr()
	if c != Nil {
		t.Errorf("expected 'nil' got '%s'", c.Repr())
	}
}

func TestString(t *testing.T) {
    v := NewString("hello")
    t.Run("value", test.NewTestHelper(test.IsValue)(v))
    t.Run("string", test.NewTestHelper(IsString)(v))
    t.Run("carcdr", test.NewTestHelper(test.IsCarCdr)(v))
    t.Run("len", test.NewTestHelper(test.IsLen)(v))
    t.Run("cons", test.NewTestHelper(test.IsCons)(v))
}
