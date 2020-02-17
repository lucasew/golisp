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
    LispValue
    Len() int
    Get(k int) LispValue
    //Set(k int, v LispValue)
}

type LispMap interface {
    LispValue
    Len() int
    Get(k LispValue) LispValue
    Set(k LispValue, v LispValue) LispValue
    Keys() LispValue // Cons
    Values() LispValue // Cons tbm
    Tuples() LispValue // Cons de cons
}

type LispValue interface {
    IsNil() bool
    Car() LispValue
    Cdr() LispValue
    Repr() string
}

type IntoLispValue interface {
    ToLispValue() LispValue
}

type LispFunction interface {
    LispValue
    LispCall(LispValue) (LispValue, error)
    IsFunctionNative() bool
}
