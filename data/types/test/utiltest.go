package test

import (
	"github.com/lucasew/golisp/data"
	"github.com/lucasew/golisp/data/entity/register"
)

func init() {
	register.Register("lisp_number", func(v data.LispValue) bool {
		_, ok := v.(data.LispNumber)
		return ok
	})
	register.Register("lisp_string", func(v data.LispValue) bool {
		_, ok := v.(data.LispString)
		return ok
	})
	register.Register("lisp_atom", func(v data.LispValue) bool {
		_, ok := v.(data.LispAtom)
		return ok
	})
	register.Register("lisp_cons", func(v data.LispValue) bool {
		_, ok := v.(data.LispCons)
		return ok
	})
	register.Register("lisp_namespace", func(v data.LispValue) bool {
		_, ok := v.(data.LispNamespace)
		return ok
	})
	register.Register("lisp_map", func(v data.LispValue) bool {
		_, ok := v.(data.LispMap)
		return ok
	})
	register.Register("lisp_value", func(v data.LispValue) bool {
		return true
	})
	register.Register("lisp_function", func(v data.LispValue) bool {
		_, ok := v.(data.LispFunction)
		return ok
	})
	register.Register("lisp_portal", func(v data.LispValue) bool {
		_, ok := v.(data.LispPortal)
		return ok
	})
	register.Register("lisp_len", func(v data.LispValue) bool {
		_, ok := v.(data.LispLen)
		return ok
	})
	register.Register("lisp_carcdr", func(v data.LispValue) bool {
		_, ok := v.(data.LispCarCdr)
		return ok
	})
	register.Register("lisp_iterator", func(v data.LispValue) bool {
		_, ok := v.(data.LispIterator)
		return ok
	})
}
