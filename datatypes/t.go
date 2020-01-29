package datatypes

type _t uint8

var T LispValue = _t(0)

func (_t) IsNil() bool {
    return false
}

func (_t) Car() LispValue {
    return Nil
}

func (_t) Cdr() LispValue {
    return Nil
}

func (_t) Repr() string {
    return "t"
}
