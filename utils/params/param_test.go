package params

import (
    "testing"
    "github.com/lucasew/golisp/data"
    "github.com/lucasew/golisp/data/types"
    "github.com/lucasew/golisp/data/types/number"
)

func TestParameterObjectLookup(t *testing.T) {
    res := NewParameterLookup(types.NewAtom("nome"), types.NewString("Lucas"), types.NewAtom("idade"), number.NewIntFromInt(19))
    if len(res.Args) != 0 {
        t.Errorf("positional Args should be zero")
    }
    if res.KwArgs["nome"].(data.LispString).ToString() != "Lucas" {
        t.Errorf("key error: nome")
    }
    n, _ := res.KwArgs["idade"].(number.LispInt).Int64()
    if n != 19 {
        t.Errorf("key error: nome")
    }
}
