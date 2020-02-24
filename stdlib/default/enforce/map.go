package enforce

import (
	"fmt"
	"github.com/lucasew/golisp/data"
	"github.com/lucasew/golisp/data/types"
)

func Map(d data.LispValue, nth int) error {
	if !types.IsMap(d) {
		return fmt.Errorf("%d nth parameter expects a map, got %s", nth, d.LispTypeName())
	}
	return nil
}