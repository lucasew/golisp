package datatypes

import (
    "fmt"
)

type KV []LispValue

func NewKV(k LispValue, v LispValue) LispValue {
    return KV([]LispValue{k, v})
}

func (k KV) Car() LispValue {
    return k[0]
}

func (k KV) Cdr() LispValue {
    return k[1]
}

func (kv KV) IsNil() bool {
    return kv[1].IsNil()
}

func (kv KV) Repr() string {
    return fmt.Sprintf("(%s : %s)", kv[0], kv[1])
}
