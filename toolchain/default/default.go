package tdefault

import (
    "github.com/lucasew/golisp/toolchain"
    "github.com/lucasew/golisp/parser/default"
    "github.com/lucasew/golisp/vm/default"
    "github.com/lucasew/golisp/vm/components/env"
)

func NewDefaultToolchain(e *env.LispEnv) toolchain.Toolchain {
    return toolchain.NewToolchain(vm_default.NewVM(e), pdefault.Parse)
}
