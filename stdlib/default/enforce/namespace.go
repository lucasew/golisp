
package enforce

import (
	"fmt"
	"github.com/lucasew/golisp/data"
	"github.com/lucasew/golisp/data/types/test"
)

func Namespace(d data.LispValue, nth int) error {
	if !test.IsNamespace(d) {
		return fmt.Errorf("%d nth parameter expects a namespace, got %s", nth, d.LispTypeName())
	}
	return nil
}
