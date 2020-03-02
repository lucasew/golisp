package enforce

import (
	"fmt"
	"github.com/lucasew/golisp/data"
)

var ErrUnmatchedLength = fmt.Errorf("unmatched length")

func Length(lst []data.LispValue, expected int) error {
	got := len(lst)
	if got != expected {
		return fmt.Errorf("%w: expected %d got %d", ErrUnmatchedLength, expected, got)
	}
	return nil
}
