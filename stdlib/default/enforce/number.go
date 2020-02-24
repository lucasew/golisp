package enforce

import (
	"fmt"
	"github.com/lucasew/golisp/data"
	"github.com/lucasew/golisp/data/types"
)

var ErrNotANumber = fmt.Errorf("not a number")

func Number(v data.LispValue) error {
	if !types.IsNumber(v) {
		return ErrNotANumber
	}
	return nil
}
