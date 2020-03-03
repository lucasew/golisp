package vm_default

import (
	// "github.com/lucasew/golisp/datatypes"
	"github.com/davecgh/go-spew/spew"
	"github.com/lucasew/golisp/data"
	parser "github.com/lucasew/golisp/parser/default"
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

func TestMap(t *testing.T) {
	ast, err := parser.Parse(`
    (setg incr (lambda (x) (+ x 1)))
    (collect (map incr '(1 2 3)))
    `)
	if err != nil {
		t.Error(err)
	}
	vm := NewVM(nil)
	ret, err := vm.Eval(ast)
	if err != nil {
		t.Error(err)
	}
	expected := " (2 3 4) "
	got := ret.Repr()
	if expected != got {
		t.Errorf("expected: %s got: %s", expected, got)
	}
}
