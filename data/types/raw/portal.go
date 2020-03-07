package raw

import (
	"github.com/lucasew/golisp/data"
	"github.com/lucasew/golisp/data/types"
	"reflect"
)

func NewPortalFromNativeChan(v interface{}) data.LispPortal {
	if reflect.TypeOf(v).Kind() == reflect.Chan {
		p := types.NewPortal(0)
		go func() {
			v := reflect.ValueOf(v)
			val, ok := v.Recv()
			if !ok {
				p.Close()
				return
			}
			p.Send(NewLispWrapper(val.Interface()))
		}()
		return p
	}
	return nil
}
