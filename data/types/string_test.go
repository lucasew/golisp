package types

import (
    "testing"
    "github.com/lucasew/golisp/data"
)

func TestStringCar(t *testing.T) {
    s := NewConventionalString("hello")
    c := s.Car()
    if c != NewByte('h') {
        t.Errorf("expected 'h' got '%c'", c)
    }
}

func TestStringCdr(t *testing.T) {
    s := NewConventionalString("hello")
    c := s.Cdr()
    if c != NewConventionalString("ello") {
        t.Errorf("expected 'hello' got '%s'", c)
    }
    c = c.Cdr()
    if c != NewConventionalString("llo") {
        t.Errorf("expected 'hello' got '%s'", c)
    }
}

func TestStringCdrSmall(t *testing.T) {
    // let x be a string that (< (len x) 2)
    // -- (assert-true (eq (cdr x) nil))
    s := NewConventionalString("a") 
    c := s.Cdr()
    if c != data.Nil {
        t.Errorf("expected 'nil' got '%s'", c.Repr())
    }
}
