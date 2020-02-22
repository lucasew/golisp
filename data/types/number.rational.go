package types

import (
    "math/big"
    "testing"
    "github.com/lucasew/golisp/data"
)

type LispRational struct {
    n *big.Rat
}

func NewRationalFromInt(n LispInt) LispRational {
    return NewRationalFromInts(n, NewIntFromInt64(1))
}

func NewRationalFromInts(num LispInt, denom LispInt) LispRational {
    r := &big.Rat{}
    r.SetFrac(num.n, denom.n)
    return NewRationalFromBigRat(r)
}

func NewRationalFromBigRat(n *big.Rat) LispRational {
    return LispRational{
        n: n,
    }
}

func NewRationalFromString(s string) (LispRational, bool) {
    ret := &big.Rat{}
    _, ok := ret.SetString(s)
    return NewRationalFromBigRat(ret), ok
}

func (i LispRational) IsZero() bool {
    return i.n.Num().Cmp(big.NewInt(0)) == 0
}

func (i LispRational) IsInt() bool {
    return i.n.IsInt() // Se denominador é 1
}

func (i LispRational) IsInfinite() bool {
    return false
}

func (i LispRational) Neg() LispRational {
    ret := &big.Rat{}
    ret.Neg(i.n)
    return NewRationalFromBigRat(ret)
}

func (i LispRational) Abs() LispRational {
    ret := &big.Rat{}
    ret.Abs(i.n)
    return NewRationalFromBigRat(ret)
}

func (i LispRational) Cmp(n LispRational) int {
    return i.n.Cmp(n.n)
}

func (i LispRational) NumberAbsCmp(n LispRational) int {
    return i.Abs().Cmp(n.Abs())
}

func (i LispRational) Sub(n LispRational) LispRational {
    ret := &big.Rat{}
    return NewRationalFromBigRat(ret.Sub(i.n, n.n))
}

func (i LispRational) Sum(n LispRational) LispRational {
    ret := &big.Rat{}
    return NewRationalFromBigRat(ret.Add(i.n, n.n))
}

func (i LispRational) Mul(n LispRational) LispRational {
    ret := &big.Rat{}
    return NewRationalFromBigRat(ret.Mul(i.n, n.n))
}

func (i LispRational) Div(n LispRational) LispRational {
    y := &big.Rat{}
    y.Inv(n.n)
    ret := &big.Rat{}
    return NewRationalFromBigRat(ret.Mul(i.n, y))
}

func (i LispRational) IsNil() bool {
    return i.n == nil
}

func (i LispRational) Repr() string {
    return i.n.RatString()
}

func (LispRational) LispTypeName() string {
    return "rational"
}

func IsRational(v data.LispValue) bool {
    _, ok := v.(LispRational)
    return ok
}

func RationalTest(v data.LispValue) func(*testing.T) {
    return func(t *testing.T) {
        if !IsRational(v) {
            t.Fail()
        }
    }
}
