package enforce

import (
	"fmt"
	"github.com/lucasew/golisp/data"
	eerr "github.com/lucasew/golisp/data/entity/error"
	"github.com/lucasew/golisp/data/entity/register"
)

func Entity(name string, d []data.LispValue, nth int) func() error {
	e, ok := register.Get(name)
	return func() error {
		if !ok {
			return fmt.Errorf("entity not found")
		}
		v := d[nth-1]
		err := eerr.EntityAsError(e)(v)
		if err != nil {
			return fmt.Errorf("%dth parameter: %w", nth, err)
		}
		return nil
	}
}
