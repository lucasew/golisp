package enforce

import (
	"fmt"
	"github.com/lucasew/golisp/data"
	"github.com/lucasew/golisp/data/types"
)

func Function(d data.LispValue, nth int) error {
	if !types.IsFunction(d) {
		return fmt.Errorf("%d nth parameter expects a function, got %s", nth, d.LispTypeName())
	}
	return nil
}
