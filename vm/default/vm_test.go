package vm_default

import (
    // "github.com/lucasew/golisp/datatypes"
    parser "github.com/lucasew/golisp/parser/default"
    "github.com/davecgh/go-spew/spew"
    "testing"
)

func TestBasicEval(t *testing.T) {
    ast, err := parser.Parse("(and 2 3 4)")
    if err != nil {
        t.Error(err)
    }
    vm := NewVM(nil)
    res, err := vm.Eval(ast)
    if err != nil {
        t.Error(err)
    }
    if res.Repr() != " (4) " {
        println(res.Repr())
        spew.Dump(res)
        t.Fail()
    }
}

func TestIfPrints(t *testing.T) {
    ast, err := parser.Parse("(println \"teste\" )")
    if err != nil {
        t.Error(err)
    }
    vm := NewVM(nil)
    res, err := vm.Eval(ast)
    if err != nil {
        t.Error(err)
    }
    if res.Repr() != " (\"teste\") " {
        t.Fail()
    }
}

func TestAcessFunctionAnd(t *testing.T) {
    vm := NewVM(nil)
    if vm.EnvGet("and").IsNil() {
        t.Fail()
    }
}
