package iterator

import (
	"github.com/lucasew/golisp/data"
	"github.com/lucasew/golisp/data/types"
    "fmt"
)

func NewIterator(v data.LispValue) (data.LispIterator, error) {
    switch in := v.(type) {
    case data.LispCons:
        return NewConsIterator(in), nil
    case data.LispIterator:
        return in, nil
    }
    return NewConsIterator(types.NewCons()), fmt.Errorf("cant convert type %s to iterator", v.LispTypeName())
}
