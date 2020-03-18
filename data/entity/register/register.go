package register

import (
	"fmt"
	"github.com/lucasew/golisp/data"
	"github.com/lucasew/golisp/data/entity"
)

var Registry = map[string]data.LispEntity{}

func Register(e data.LispEntity) {
	_, ok := Registry[e.EntityName()]
	if ok {
		err := fmt.Errorf("collision: cant register %s entity twice", e.EntityName())
		panic(err)
	}
	Registry[e.EntityName()] = e
}

func BuildAndRegister(name string, isfunc func(v data.LispValue) bool) {
	Register(entity.Entity{
		Name: name,
		Isfn: isfunc,
	})
}

func Get(k string) (data.LispEntity, bool) {
	r, ok := Registry[k]
	return r, ok
}

func Is(k string, v data.LispValue) bool {
	e, ok := Get(k)
	if !ok {
		return ok
	}
	return e.EntityIsFn(v)
}
