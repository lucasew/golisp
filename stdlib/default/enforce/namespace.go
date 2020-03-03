package enforce

import (
	"fmt"
	"github.com/lucasew/golisp/data"
	"github.com/lucasew/golisp/data/types/test"
)

func Namespace(d []data.LispValue, nth int) func() error {
	return func() error {
		v := d[nth-1]
		if !test.IsNamespace(v) {
			return fmt.Errorf("%d nth parameter expects a namespace, got %s", nth, v.LispTypeName())
		}
		return nil
	}
}
