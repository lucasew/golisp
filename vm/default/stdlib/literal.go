package stdlib

import (
    "github.com/lucasew/golisp/data"
    "github.com/lucasew/golisp/vm"
)

func init() {
    register("quote", Quote)
    register("list", List)
}

func Quote(env vm.LispVM, v data.LispValue) (data.LispValue, error) {
    return v.Car(), nil
}

func List(env vm.LispVM, v data.LispValue) (data.LispValue, error) {
    return v, nil
}
