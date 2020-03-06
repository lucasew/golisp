package time

import (
	"github.com/lucasew/golisp/data/types/test"
	"testing"
    "time"
)

func TestTimeType(t *testing.T) {
	v := NewTimeFromTime(time.Now())
	t.Run("lisp_value", test.NewTestHelper(test.IsValue)(v))
	t.Run("map", test.NewTestHelper(test.IsMap)(v))
}
