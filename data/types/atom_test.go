package types

import (
	"testing"
)

func TestAtomType(t *testing.T) {
	t.Run("lisp_value", ValueTest(NewAtom("")))
	t.Run("atom", AtomTest(NewAtom("")))
}

var AtomTest = NewTestHelper(IsAtom)
