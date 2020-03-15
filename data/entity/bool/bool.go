package bool

import (
	"github.com/lucasew/golisp/data"
	"github.com/lucasew/golisp/data/entity"
)

func EntityAsBool(e entity.Entity) func(v data.LispValue) bool {
	return e.Isfn
}
