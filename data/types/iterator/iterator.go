package iterator

import (
	"github.com/lucasew/golisp/data"
)

func IsIterator(v data.LispValue) bool {
	_, ok := v.(data.LispIterator)
	return ok
}
