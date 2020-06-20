package enforce

import (
	"fmt"
	"github.com/lucasew/golisp/data"
	"github.com/lucasew/golisp/data/types"
)

func Portal(d []data.LispValue, nth int) func() error {
	return func() error {
		v := d[nth-1]
		if !types.IsPortal(v) {
			return fmt.Errorf("%d nth parameter expects a portal, got %s", nth, v.LispTypeName())
		}
		return nil
	}
}
