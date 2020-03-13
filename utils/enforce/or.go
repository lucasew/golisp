package enforce 

import (
    "fmt"
)

func Or(conds ...func()error ) func() error {
    return func () (err error) {
        for _, cond := range conds {
            err = cond()
            if err == nil {
                return
            }
            err = fmt.Errorf("or: %s", err.Error())
        }
        return
    }
}
