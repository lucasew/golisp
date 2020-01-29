package datatypes

import (
    "math/big"
)

type LispInt struct {
    n *big.Int
}

func NewIntFromInt64(n int64) LispInt {
    return LispInt{
        n: big.NewInt(n),
    }
}

func NewIntFromBigInt(n *big.Int) LispInt {
    return LispInt{
        n: n,
    }
}

func (n LispInt) Int64() (r int64, acc LispAccuracy) {
    if (n.n.IsInt64()) {
        acc = AccuracyExact
    } else {
        acc = AccuracyInvalid
    }
    r = n.n.Int64()
    return
}

func (n LispInt) Uint64() (r uint64, acc LispAccuracy) {
    if (n.n.IsUint64()) {
        acc = AccuracyExact
    } else {
        acc = AccuracyInvalid
    }
    r = n.n.Uint64()
    return
}

func (n LispInt) IsZero() bool {
    return n.n.Cmp(big.NewInt(0)) == 0
}

func (LispInt) IsInt() bool {
    return true
}

func (LispInt) IsInfinite() bool {
    return false
}

func (n LispInt) Neg() LispInt {
    num := big.NewInt(0)
    num.Neg(n.n)
    return NewIntFromBigInt(num)
}

func (n LispInt) Cmp(other LispInt) int {
    return n.n.Cmp(other.n)
}

func (n LispInt) AbsCmp(other LispInt) int {
    return n.n.CmpAbs(other.n)
}

func (n LispInt) Sub(other LispInt) LispInt {
    num := &big.Int{}
    return NewIntFromBigInt(num.Sub(n.n, other.n))
}

func (n LispInt) Sum(other LispInt) LispInt {
    num := &big.Int{}
    return NewIntFromBigInt(num.Add(n.n, other.n))
}

func (n LispInt) Mul(other LispInt) LispInt {
    num := &big.Int{}
    return NewIntFromBigInt(num.Mul(n.n, other.n))
}

func (n LispInt) Div(other LispInt) LispInt {
    num := &big.Int{}
    return NewIntFromBigInt(num.Div(n.n, other.n))
}

func (n LispInt) Mod(other LispInt) LispInt {
    num := &big.Int{}
    return NewIntFromBigInt(num.Mod(n.n, other.n))
}

func (n LispInt) Rem(other LispInt) LispInt {
    num := &big.Int{}
    return NewIntFromBigInt(num.Rem(n.n, other.n))
}

// QuoRem returns a/b and a%b
func (n LispInt) QuoRem(other LispInt, mod LispInt) (LispInt, LispInt) {
    qi := &big.Int{}
    ri := &big.Int{}
    qi, ri = qi.QuoRem(n.n, other.n, ri)
    return NewIntFromBigInt(qi), NewIntFromBigInt(ri)
}

func (n LispInt) Sqrt() LispInt {
    ret := &big.Int{}
    return NewIntFromBigInt(ret.Sqrt(n.n))
}

func (n LispInt) Exp(other LispInt) LispInt {
    ret := &big.Int{}
    return NewIntFromBigInt(ret.Exp(n.n, other.n, nil))
}

func (n LispInt) Abs() LispInt {
    num := &big.Int{}
    num.Abs(n.n)
    return NewIntFromBigInt(num)
}

func (LispInt) Car() LispValue {
    return Nil
}

func (LispInt) Cdr() LispValue {
    return Nil
}

func (i LispInt) IsNil() bool {
    return i.n == nil
}

func (i LispInt) Repr() string {
    return i.n.String()
}
