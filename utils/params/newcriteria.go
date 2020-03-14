package params

import (
    "github.com/lucasew/golisp/data"
    "fmt"
)

func NewCriteriaForIsFuncion(isfn func(data.LispValue)bool, name string) func(data.LispValue) error {
    return func(v data.LispValue) error {
        if isfn(v) {
            return nil
        }
        return fmt.Errorf("value does not pass the criteria %s", name)
    }
}
