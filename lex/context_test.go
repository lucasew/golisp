package lex

import (
    "testing"
)

func TestShouldIncrementTheCounter(t *testing.T) {
    ctx := Context{
        data: []byte("teste eoq trabson"),
        index: 0,
    }
    if ctx.Index() != 0 {
        t.Errorf("counter needs to start at 0")
    }
    ctx.Increment()
    if ctx.Index() != 1 {
        t.Errorf("not incremented")
    }
}

