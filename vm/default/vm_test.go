package vm_default

import (
    // "github.com/lucasew/golisp/datatypes"
    parser "github.com/lucasew/golisp/parser/default"
    "github.com/lucasew/golisp/data"
    "github.com/davecgh/go-spew/spew"
    "testing"
)

func TestBasicEval(t *testing.T) {
    ast, err := parser.Parse("(and (or 2 3 4) 3 4)")
    if err != nil {
        t.Error(err)
    }
    vm := NewVM(nil)
    res, err := vm.Eval(ast)
    if err != nil {
        t.Error(err)
    }
    if res.Repr() != "4" {
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
    if res.Repr() != "\"teste\"" {
        t.Fail()
    }
}

func TestAcessFunctionAnd(t *testing.T) {
    vm := NewVM(nil)
    if vm.EnvGet("and").IsNil() {
        t.Fail()
    }
}

func TestSetg(t *testing.T) {
    ast, err := parser.Parse("(setg teste \"teste\")")
    if err != nil {
        t.Error(err)
    }
    vm := NewVM(nil)
    _, err = vm.Eval(ast)
    if err != nil {
        t.Error(err)
    }
    r := vm.EnvGet("teste")
    s := r.(data.LispString)
    if s.ToString() != "teste" {
        t.Fail()
    }

}

func TestIfYes(t *testing.T) {
    ast, err := parser.Parse("(if 2 3 4)") // 2 == true
    if err != nil {
        t.Error(err)
    }
    vm := NewVM(nil)
    ret, err := vm.Eval(ast)
    if err != nil {
        t.Error(err)
    }
    expected := "3"
    got := ret.Repr()
    if expected != got {
        t.Error("expected: ", expected, " got: ", got)
    }
}

func TestIfNo(t *testing.T) {
    ast, err := parser.Parse("(if nil 3 4)") // nil == false
    if err != nil {
        t.Error(err)
    }
    vm := NewVM(nil)
    ret, err := vm.Eval(ast)
    if err != nil {
        t.Error(err)
    }
    expected := "4"
    got := ret.Repr()
    if expected != got {
        t.Error("expected: ", expected, " got: ", got)
    }
}

