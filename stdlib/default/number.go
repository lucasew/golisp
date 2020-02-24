package stdlib

import (
	"fmt"
	"github.com/lucasew/golisp/data"
	"github.com/lucasew/golisp/data/types"
	"github.com/lucasew/golisp/data/types/number"
	"github.com/lucasew/golisp/stdlib/default/enforce"
)

func init() {
	register("to-float", ToFloat)
	register("to-rat", ToRat)
	register("to-int", ToInt)
	register("to-byte", ToByte)
}

func ToFloat(v data.LispCons) (data.LispValue, error) {
	switch n := v.Car().(type) {
	case number.LispFloat:
		return n, nil
	case number.LispInt:
		return number.NewFloatFromInt(n), nil
	case number.LispRational:
		return number.NewFloatFromRational(n), nil
	default:
		return types.Nil, fmt.Errorf("cant convert %T to float", n)
	}
}

func ToRat(v data.LispCons) (data.LispValue, error) {
	switch n := v.Car().(type) {
	case number.LispFloat:
		r, ok := number.NewRationalFromString(n.Repr())
		if !ok {
			panic("invalid state: cant parse rational")
		}
		return r, nil
	case number.LispInt:
		return number.NewRationalFromInt(n), nil
	case number.LispRational:
		return n, nil
	default:
		return types.Nil, fmt.Errorf("cant convert %T to rational", n)
	}
}

func ToInt(v data.LispCons) (data.LispValue, error) {
	switch n := v.Car().(type) {
	case number.LispFloat:
		return number.NewIntFromFloat(n)
	case number.LispRational:
		return number.NewIntFromFloat(number.NewFloatFromRational(n))
	case number.LispInt:
		return n, nil
	default:
		return types.Nil, fmt.Errorf("cant convert %T to int", n)
	}
}

func ToByte(v data.LispCons) (data.LispValue, error) {
	err := enforce.Validate(enforce.Length(v, 1), enforce.Number(v.Car()))
	if err != nil {
		return types.Nil, err
	}
	num := number.NewByte(0)
	vnum := v.Car().(data.LispNumber)
	switch n := vnum.(type) {
	case number.LispInt:
		tmp, _ := n.Int64()
		num = number.NewByte(byte(tmp))
	case number.Byte:
		num = n
	default:
		return types.Nil, fmt.Errorf("invalid type for the first parameter, expected byte got %s", vnum.LispTypeName())
	}
	return num, nil
}
