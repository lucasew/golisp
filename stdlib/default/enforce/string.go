package enforce

import (
    "fmt"
    "github.com/lucasew/golisp/data"
    "github.com/lucasew/golisp/data/types"
)

func String(d []data.LispValue, nth int) func()error {
    return func()error {
        v := d[nth - 1]
        if !types.IsString(v) {
            return fmt.Errorf("%d nth element is not a string", nth)
        }
        return nil
    }
}
