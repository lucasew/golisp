package data

import (
	"context"
	"errors"
)

var ErrContextCancelled = errors.New("context canceled")

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
	LispEntity() LispEntity
	Repr() string
}

type LispEntity interface {
	EntityName() string
	EntityIsFn(LispValue) bool
	AssignTo(LispValue, interface{}) bool
	EntitySubEntities() LispSubEntities
}

type LispSubEntities interface {
	LookupSubEntity(k string) LispSubEntity
	Keys() []string
	AddSubEntity(k string, s LispSubEntity) LispSubEntities
}

type LispSubEntity interface {
	Get(self LispValue) LispValue
	Set(self LispValue, to LispValue) (outself LispValue)
}

type IntoLispValue interface {
	ToLispValue() LispValue
}

type LispFunction interface {
	LispValue
	LispCall(context.Context, ...LispValue) (LispValue, error)
	IsFunctionNative() bool
}

type LispPortal interface {
	LispValue
	Send(v LispValue) LispValue
	SendUnblocking(v LispValue) LispValue
	Recv() LispValue
	RecvUnblocking() LispValue
	Close()
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
	Next(context.Context) LispValue
	IsEnd(context.Context) bool
}
