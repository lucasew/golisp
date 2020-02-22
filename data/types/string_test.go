package types

import (
    "testing"
)

func TestStringCar(t *testing.T) {
    s := NewString("hello")
    c := s.Car()
    if c != NewByte('h') {
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
