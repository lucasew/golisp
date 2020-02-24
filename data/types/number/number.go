package number

import (
	"github.com/lucasew/golisp/data"
)

func IsNumber(v data.LispValue) bool {
	_, ok := v.(data.LispNumber)
	return ok
}
