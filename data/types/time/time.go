package time

import (
	"github.com/lucasew/golisp/data"
	"github.com/lucasew/golisp/data/entity/register"
	"github.com/lucasew/golisp/data/types"
	"github.com/lucasew/golisp/data/types/maps"
	"github.com/lucasew/golisp/data/types/number"
	"time"
)

type Time time.Time

func init() {
	register.Register("time", func(v data.LispValue) bool {
		_, ok := v.(Time)
		return ok
	})
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
	return t.Unwrap().String()
}

func (Time) LispTypeName() string {
	return "time"
}

func (t Time) Unwrap() time.Time {
	return time.Time(t)
}

func (t Time) Get(k data.LispValue) data.LispValue {
	return maps.MappeableGet(t, k)
}

func (t Time) Set(k data.LispValue, v data.LispValue) data.LispValue {
	return maps.MappeableSet(t, k, v)
}

func (t Time) Keys() data.LispCons {
	return maps.MappeableKeys(t)
}

func (t Time) Values() data.LispCons {
	return maps.MappeableValues(t)
}

func (t Time) Tuples() data.LispCons {
	return maps.MappeableTuples(t)
}

func (t Time) Len() int {
	return maps.MappeableLen(t)
}

func (t Time) GettableElements() map[string]func() data.LispValue {
	return map[string]func() data.LispValue{
		"day": func() data.LispValue {
			return number.NewIntFromInt(t.Unwrap().Day())
		},
		"month_int": func() data.LispValue {
			return number.NewIntFromInt(int(t.Unwrap().Month()))
		},
		"month_string": func() data.LispValue {
			return types.NewString(t.Unwrap().Month().String())
		},
		"year": func() data.LispValue {
			return number.NewIntFromInt(t.Unwrap().Year())
		},
		"year_day": func() data.LispValue {
			return number.NewIntFromInt(t.Unwrap().YearDay())
		},
		"hour": func() data.LispValue {
			return number.NewIntFromInt(t.Unwrap().Hour())
		},
		"minute": func() data.LispValue {
			return number.NewIntFromInt(t.Unwrap().Minute())
		},
		"second": func() data.LispValue {
			return number.NewIntFromInt(t.Unwrap().Second())
		},
		"nanosecond": func() data.LispValue {
			return number.NewIntFromInt(t.Unwrap().Nanosecond())
		},
		"unix": func() data.LispValue {
			return number.NewIntFromInt64(t.Unwrap().Unix())
		},
	}
}

func (t Time) SettableElements() map[string]func(data.LispValue) data.LispValue {
	return map[string]func(data.LispValue) data.LispValue{}
}
