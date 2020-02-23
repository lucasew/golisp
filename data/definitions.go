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
	LispCons
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
	//Set(k int, v LispValue)
}

type LispMap interface {
	LispValue
	LispLen
	Get(k LispValue) LispValue
	Set(k LispValue, v LispValue) LispValue
	Keys() LispValue   // Cons
	Values() LispValue // Cons tbm
	Tuples() LispValue // Cons de cons
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
	LispCall(LispCons) (LispValue, error)
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
