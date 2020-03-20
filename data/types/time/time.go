package time

import (
	"github.com/lucasew/golisp/data"
	"github.com/lucasew/golisp/data/entity"
	"github.com/lucasew/golisp/data/entity/register"
	"time"
)

type Time time.Time

func init() {
	register.Register(new(Time).LispEntity())
}

func (Time) LispEntity() data.LispEntity {
	return entity.Entity{
		"time", func(v data.LispValue) bool {
			_, ok := v.(Time)
			return ok
		},
	}
}

func NewTimeFromUnix(n int64) Time {
	return Time(time.Unix(n, 0))
}

func NewTimeFromTime(t time.Time) Time {
	return Time(t)
}

func (t Time) IsNil() bool {
	return false
}

func (t Time) Repr() string {
	return time.Time(t).String()
}
