package test

import (
	"github.com/lucasew/golisp/data"
	"github.com/lucasew/golisp/data/entity/register"
)

func init() {
	register.BuildAndRegister("lisp_number", func(v data.LispValue) bool {
		_, ok := v.(data.LispNumber)
		return ok
	})
	register.BuildAndRegister("lisp_string", func(v data.LispValue) bool {
		_, ok := v.(data.LispString)
		return ok
	})
	register.BuildAndRegister("lisp_atom", func(v data.LispValue) bool {
		_, ok := v.(data.LispAtom)
		return ok
	})
	register.BuildAndRegister("lisp_cons", func(v data.LispValue) bool {
		_, ok := v.(data.LispCons)
		return ok
	})
	register.BuildAndRegister("lisp_namespace", func(v data.LispValue) bool {
		_, ok := v.(data.LispNamespace)
		return ok
	})
	register.BuildAndRegister("lisp_map", func(v data.LispValue) bool {
		_, ok := v.(data.LispMap)
		return ok
	})
	register.BuildAndRegister("lisp_value", func(v data.LispValue) bool {
		return true
	})
	register.BuildAndRegister("lisp_function", func(v data.LispValue) bool {
		_, ok := v.(data.LispFunction)
		return ok
	})
	register.BuildAndRegister("lisp_portal", func(v data.LispValue) bool {
		_, ok := v.(data.LispPortal)
		return ok
	})
	register.BuildAndRegister("lisp_len", func(v data.LispValue) bool {
		_, ok := v.(data.LispLen)
		return ok
	})
	register.BuildAndRegister("lisp_carcdr", func(v data.LispValue) bool {
		_, ok := v.(data.LispCarCdr)
		return ok
	})
	register.BuildAndRegister("lisp_iterator", func(v data.LispValue) bool {
		_, ok := v.(data.LispIterator)
		return ok
	})
}
