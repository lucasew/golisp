package stdlib

import (
	"fmt"
	"github.com/lucasew/golisp/data"
	"github.com/lucasew/golisp/data/types"
	"github.com/lucasew/golisp/vm/default/stdlib/enforce"
)

func init() {
	register("to-float", ToFloat)
	register("to-rat", ToRat)
	register("to-int", ToInt)
	register("to-byte", ToByte)
}

func ToFloat(v data.LispCons) (data.LispValue, error) {
	switch n := v.Car().(type) {
	case types.LispFloat:
		return n, nil
	case types.LispInt:
		return types.NewFloatFromInt(n), nil
	case types.LispRational:
		return types.NewFloatFromRational(n), nil
	default:
		return types.Nil, fmt.Errorf("cant convert %T to float", n)
	}
}

func ToRat(v data.LispCons) (data.LispValue, error) {
	switch n := v.Car().(type) {
	case types.LispFloat:
		r, ok := types.NewRationalFromString(n.Repr())
		if !ok {
			panic("invalid state: cant parse rational")
		}
		return r, nil
	case types.LispInt:
		return types.NewRationalFromInt(n), nil
	case types.LispRational:
		return n, nil
	default:
		return types.Nil, fmt.Errorf("cant convert %T to rational", n)
	}
}

func ToInt(v data.LispCons) (data.LispValue, error) {
	switch n := v.Car().(type) {
	case types.LispFloat:
		return types.NewIntFromFloat(n)
	case types.LispRational:
		return types.NewIntFromFloat(types.NewFloatFromRational(n))
	case types.LispInt:
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
	number := types.NewByte(0)
	vnum := v.Car().(data.LispNumber)
	switch n := vnum.(type) {
	case types.LispInt:
		tmp, _ := n.Int64()
		number = types.NewByte(byte(tmp))
	case types.Byte:
		number = n
	default:
		return types.Nil, fmt.Errorf("invalid type for the first parameter, expected byte got %s", vnum.LispTypeName())
	}
	return number, nil
}
