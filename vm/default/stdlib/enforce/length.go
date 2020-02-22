package enforce

import (
    "github.com/lucasew/golisp/data"
    "fmt"
)

var ErrUnmatchedLength = fmt.Errorf("unmatched length")

func Length(lst data.LispCons, expected int) error {
    got := lst.Len()
    if got != expected {
        return fmt.Errorf("%w: expected %d got %d", ErrUnmatchedLength, expected, got)
    }
    return nil
}
