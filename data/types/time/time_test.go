package time

import (
	"github.com/lucasew/golisp/data/entity/test"
	"testing"
	"time"
)

func TestTimeType(t *testing.T) {
	v := NewTimeFromTime(time.Now())
	test.TestValues(v, t, "time")
}
