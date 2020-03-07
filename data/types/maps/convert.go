package maps

import (
	"fmt"
	"github.com/lucasew/golisp/data"
)

func NewMapFromInterface(v interface{}) (data.LispMap, error) {
	switch mi := v.(type) {
	case data.LispMap:
		return mi, nil
	case map[string]data.LispValue:
		return NewMapFromMapString(mi), nil
	case map[data.LispValue]data.LispValue:
		return NewMapFromMapValue(mi), nil
	case Mappeable:
		return NewMapFromMappeable(mi), nil
	default:
		return nil, fmt.Errorf("invalid type: %T", mi)
	}
}
