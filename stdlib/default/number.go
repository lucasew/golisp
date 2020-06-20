package stdlib

import (
	"context"
	"fmt"
	"github.com/lucasew/golisp/data"
	"github.com/lucasew/golisp/data/types"
	"github.com/lucasew/golisp/data/types/number"
	"github.com/lucasew/golisp/utils/enforce"
)

func init() {
	register("to-float", ToFloat)
	register("to-rat", ToRat)
	register("to-int", ToInt)
	register("to-byte", ToByte)
}

func ToFloat(ctx context.Context, v ...data.LispValue) (data.LispValue, error) {
	switch n := v[0].(type) {
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

func ToRat(ctx context.Context, v ...data.LispValue) (data.LispValue, error) {
	switch n := v[0].(type) {
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

func ToInt(ctx context.Context, v ...data.LispValue) (data.LispValue, error) {
	switch n := v[0].(type) {
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

func ToByte(ctx context.Context, v ...data.LispValue) (data.LispValue, error) {
	err := enforce.Validate(enforce.Length(v, 1), enforce.Number(v, 1))
	if err != nil {
		return types.Nil, err
	}
	num := number.NewByte(0)
	vnum := v[0].(data.LispNumber)
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
