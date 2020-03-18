package error

import (
	"fmt"
	"github.com/lucasew/golisp/data"
)

func EntityAsError(e data.LispEntity) func(data.LispValue) error {
	return func(v data.LispValue) error {
		if !e.EntityIsFn(v) {
			return fmt.Errorf("expected type that belongs %s got %s", e.EntityName(), v.LispEntity().EntityName())
		}
		return nil
	}
}
