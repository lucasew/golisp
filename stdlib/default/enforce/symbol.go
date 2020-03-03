package enforce

import (
    "fmt"
    "github.com/lucasew/golisp/data"
    "github.com/lucasew/golisp/data/types"
)

func Symbol(d []data.LispValue, nth int) func()error {
    return func()error {
        v := d[nth - 1]
        if !types.IsSymbol(v) {
            return fmt.Errorf("%d nth element is not avsymbol", nth)
        }
        return nil
    }
}
