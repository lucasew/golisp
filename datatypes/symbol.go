package datatypes

type Symbol string

func NewSymbol(s string) LispValue {
    return Symbol(s)
}

func (Symbol) Car() LispValue {
    return Nil
}

func (Symbol) Cdr() LispValue {
    return Nil
}

func (Symbol) IsNil() bool {
    return false
}

func (s Symbol) Repr() string {
    return string(s)
}
