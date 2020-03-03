package enforce

import (
    "fmt"
    "github.com/lucasew/golisp/data"
    "github.com/lucasew/golisp/data/types/test"
)

func Map(d []data.LispValue, nth int) func()error {
    return func()error {
        v := d[nth - 1]
        if !test.IsMap(v) {
            return fmt.Errorf("%d nth parameter expects a map, got %s", nth, v.LispTypeName())
        }
        return nil
    }
}
