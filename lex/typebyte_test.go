package lex

import (
    "testing"
)

func TestByteNewline(t *testing.T) {
    if !LexByte('\n').IsBlank() {
        t.Fail()
    }
}

func TestByteSpace(t *testing.T) {
    if !LexByte(' ').IsBlank() {
        t.Fail()
    }
}



