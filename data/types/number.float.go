package types

import (
    "math/big"
    "github.com/lucasew/golisp/data"
)

type LispFloat struct {
    n *big.Float
}

func NewFloatFromFloat64(n float64) LispFloat {
    return NewFloatFromBigFloat(big.NewFloat(n))
}

func NewFloatFromBigFloat(n *big.Float) LispFloat {
    return LispFloat{
        n: n,
    }
}

func NewFloatFromString(s string) (LispFloat, bool) {
    ret := &big.Float{}
    _, ok := ret.SetString(s)
    return NewFloatFromBigFloat(ret), ok
}

func NewFloatFromInt(i LispInt) LispFloat {
    ret := &big.Float{}
    return NewFloatFromBigFloat(ret.SetInt(i.n))
}

func NewFloatFromRational(i LispRational) LispFloat {
    ret := &big.Float{}
    return NewFloatFromBigFloat(ret.SetRat(i.n))
}

func (n LispFloat) IsZero() bool {
    return n.n.Cmp(big.NewFloat(0)) == 0
}

func (n LispFloat) IsInt() bool {
    return n.n.IsInt()
}

func (n LispFloat) IsInfinite() bool {
    return n.n.IsInf()
}

func (n LispFloat) Neg() LispFloat {
    ret := &big.Float{}
    return NewFloatFromBigFloat(ret.Neg(n.n))
}

func (n LispFloat) Cmp(other LispFloat) int {
    return n.n.Cmp(other.n)
}

func (n LispFloat) AbsCmp(other LispFloat) int {
    return n.Abs().Cmp(other.Abs())
}

func (n LispFloat) Abs() LispFloat {
    ret := &big.Float{}
    return NewFloatFromBigFloat(ret.Abs(n.n))
}

func (n LispFloat) Sub(other LispFloat) LispFloat {
    ret := &big.Float{}
    return NewFloatFromBigFloat(ret.Sub(n.n, other.n))
}

func (n LispFloat) Sum(other LispFloat) LispFloat {
    ret := &big.Float{}
    return NewFloatFromBigFloat(ret.Add(n.n, other.n))
}


func (n LispFloat) Mul(other LispFloat) LispFloat {
    ret := &big.Float{}
    return NewFloatFromBigFloat(ret.Mul(n.n, other.n))
}

func (n LispFloat) Div(other LispFloat) LispFloat {
    ret := &big.Float{}
    return NewFloatFromBigFloat(ret.Quo(n.n, other.n))
}

func (n LispFloat) Sqrt() LispFloat {
    ret := &big.Float{}
    return NewFloatFromBigFloat(ret.Sqrt(n.n))
}

func (i LispFloat) IsNil() bool {
    return i.n == nil
}

func (i LispFloat) Repr() string {
    return i.n.String()
}

func (LispFloat) LispTypeName() string {
    return "float"
}

func IsFloat(v data.LispValue) bool {
    _, ok := v.(LispFloat)
    return ok
}

var FloatTest = NewTestHelper(IsFloat)
