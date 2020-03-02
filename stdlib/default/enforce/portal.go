package enforce

import (
	"fmt"
	"github.com/lucasew/golisp/data"
	"github.com/lucasew/golisp/data/types"
)

func Portal(v data.LispValue, nth int) error {
	if !types.IsPortal(v) {
		return fmt.Errorf("%d nth parameter expects a portal, got %s", nth, v.LispTypeName())
	}
	return nil
}
