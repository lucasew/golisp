package register

import (
	"fmt"
	"github.com/lucasew/golisp/data"
	"github.com/lucasew/golisp/data/entity"
)

var Registry = map[string]entity.Entity{}

func Register(name string, isfn func(data.LispValue) bool) {
	_, ok := Registry[name]
	if ok {
		err := fmt.Errorf("collision: cant register %s entity twice", name)
		panic(err)
	}
	e := entity.Entity{
		Name: name,
		Isfn: isfn,
	}
	Registry[name] = e
}

func Get(k string) (entity.Entity, bool) {
	r, ok := Registry[k]
	return r, ok
}

func Is(k string, v data.LispValue) bool {
	e, ok := Get(k)
	if !ok {
		return ok
	}
	return e.Isfn(v)
}
