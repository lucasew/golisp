package enforce

import (
	"fmt"
	"github.com/lucasew/golisp/data"
	"reflect"
)

var ErrNotSameType = fmt.Errorf("not the same type")

func SameType(v []data.LispValue, ntha, nthb int) func() error {
	return func() error {
		a := v[ntha-1]
		b := v[nthb-1]
		ta := reflect.TypeOf(a)
		tb := reflect.TypeOf(b)
		if ta != tb {
			return fmt.Errorf("%w: expected %s got %s for %dth and %dth value", ErrNotSameType, a.LispEntity().EntityName(), b.LispEntity().EntityName(), ntha, nthb)
		}
		return nil
	}
}
