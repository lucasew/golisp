package datatypes

type LispFunction interface {
    LispValue
    LispCall(LispValue) (LispValue, error)
    IsFunctionNative() bool
}

type lispFunction func(LispValue)(LispValue, error)

func (f lispFunction) Car() LispValue {
    return Nil
}

func (f lispFunction) Cdr() LispValue {
    return Nil
}

func (f lispFunction) IsNil() bool {
    return f == nil
}

func (f lispFunction) LispCall(i LispValue) (LispValue, error) {
    return f(i)
}

func (f lispFunction) Repr() string {
    return "<native function>"
}

func (f lispFunction) IsFunctionNative() bool {
    return true
}

func NewFunction(f func(LispValue)(LispValue, error)) LispFunction {
    return lispFunction(f)
}
