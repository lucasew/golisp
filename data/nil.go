package data

type _nil uint8

var Nil LispValue = _nil(0)

func (_nil) IsNil() bool {
    return true
}

func (_nil) Car() LispValue {
    return Nil
}

func (_nil) Cdr() LispValue {
    return Nil
}

func (_nil) Repr() string {
    return "nil"
}
