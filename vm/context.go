package vm

import (
	"github.com/lucasew/golisp/data"
)

type LispVM interface {
	Eval(data.LispValue) (data.LispValue, error)
	EnvGet(key string) data.LispValue // if not exist return nil
	EnvSetLocal(key string, value data.LispValue) data.LispValue
	EnvSetGlobal(key string, value data.LispValue) data.LispValue
}
