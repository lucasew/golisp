package pdefault

import (
	"testing"
	// "github.com/davecgh/go-spew/spew"
	"context"
)

var ctx = context.Background()

func TestGlobalState(t *testing.T) {
	//     in := `
	// (defun say-my-name (:name string)
	//     (println (concat "my name is " name)))
	//     (assert-eq (+ 2 2) 4)
	//     ))))
	// `
	/*
		3
		(eoq )
		(catapimbas )
		2
	*/
	in := `
(:atom)
(+ 2/4 2.3_3 5_)
2
(eoq)
(#x16)

("some string"

(#x20)
#_"trabzera"
)

(catapimbas) --this is a comment

"catchuga"
(defun say-name(name)
    (println "the name is " name))

`
	// in := "((trabson(((eoq)))"
	ret, err := Parse(ctx, in)
	if err != nil {
		t.Error(err)
	}
	// spew.Dump(ret)

	println(ret.Repr())
}
