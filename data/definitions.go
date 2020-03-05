package data

type LispAccuracy int

const (
	AccuracyExact = iota
	AccuracyInexact
	AccuracyAbove
	AccuracyBelow
	AccuracyInvalid
)

type LispNumber interface {
	LispValue
	IsZero() bool
	IsInt() bool
	IsInfinite() bool
}

type LispString interface {
    LispValue
	ToString() string
}

type LispAtom interface {
	LispValue
	AtomString() string
}

type LispCons interface {
	LispCarCdr
	LispLen
	Get(k int) LispValue
	UnwrapCons() ([]LispValue, error)
	//Set(k int, v LispValue)
}

type LispNamespace interface {
	LispValue
	Get(k LispValue) LispValue
	Set(k LispValue, v LispValue) LispValue
}

type LispMap interface {
	LispNamespace
	LispLen
	Keys() LispCons   // Cons
	Values() LispCons // Cons tbm
	Tuples() LispCons // Cons de cons
}

type LispValue interface {
	IsNil() bool
	LispTypeName() string
	Repr() string
}

type IntoLispValue interface {
	ToLispValue() LispValue
}

type LispFunction interface {
	LispValue
	LispCall(...LispValue) (LispValue, error)
	IsFunctionNative() bool
}

type LispPortal interface {
	LispValue
	Send(v LispValue) LispValue
	SendUnblocking(v LispValue) LispValue
	Recv() LispValue
	RecvUnblocking() LispValue
}

type LispLen interface {
	Len() int
}

type LispCarCdr interface {
	LispValue
	Car() LispValue
	Cdr() LispCarCdr
}

type LispIterator interface {
	LispValue
	Next() LispValue
	IsEnd() bool
}
