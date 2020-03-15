package error

import (
	"fmt"
	"github.com/lucasew/golisp/data"
	"github.com/lucasew/golisp/data/entity"
)

func EntityAsError(e entity.Entity) func(data.LispValue) error {
	return func(v data.LispValue) error {
		if !e.Isfn(v) {
			return fmt.Errorf("expected type that belongs %s got %s", e.Name, v.LispTypeName())
		}
		return nil
	}
}
