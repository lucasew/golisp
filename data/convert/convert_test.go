package convert

import (
    "testing"
    "github.com/lucasew/golisp/data/types"
    "reflect"
)

func TestListOfStrings(t *testing.T) {
    in := []string{
        "spam",
        "eggs",
        "foo",
        "bar",
    }
    ret, err := NewLispValue(in)
    if err != nil {
        panic(err)
    }
    got := ret.Repr()
    expected := " (\"spam\" \"eggs\" \"foo\" \"bar\") "
    if !reflect.DeepEqual(expected, got) {
        t.Errorf("expected %s got %s", expected, got)
    }
}

func TestHybridList(t *testing.T) {
    ret, err := NewLispList(2, "a", []string{"coronga", "virus"})
    if err != nil {
        t.Error(err)
    }
    got := ret.Repr()
    expected := " (2 \"a\"  (\"coronga\" \"virus\") ) "
    if !reflect.DeepEqual(expected, got) {
        t.Errorf("expected %s got %s", expected, got)
    }
}

func TestLispValueToConvert(t *testing.T) {
    in := types.NewConventionalString("teste")
    ret, err := NewLispValue(in)
    if err != nil {
        t.Error(err)
    }
    if ret != in {
        t.Errorf("LispValues should not be touched in the convert function")
    }
}
