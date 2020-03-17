package vm

import (
	"context"
	"github.com/lucasew/golisp/data"
)

type LispVM interface {
	Eval(context.Context, data.LispValue) (data.LispValue, error)
	EnvGet(key string) data.LispValue // if not exist return nil
	EnvSetLocal(key string, value data.LispValue) data.LispValue
	EnvSetGlobal(key string, value data.LispValue) data.LispValue
	Import(map[string]interface{})
	PushVM() LispVM
}
