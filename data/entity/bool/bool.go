package bool

import (
	"github.com/lucasew/golisp/data"
)

func EntityAsBool(e data.LispEntity) func(v data.LispValue) bool {
	return e.EntityIsFn
}
