package enforce

import (
	"fmt"
	"github.com/lucasew/golisp/data"
	"github.com/lucasew/golisp/data/types"
)

func Symbol(v data.LispValue, nth int) error {
	if !types.IsSymbol(v) {
		return fmt.Errorf("%d nth element is not a symbol", nth)
	}
	return nil
}
