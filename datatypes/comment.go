package datatypes

type _comment uint8

var Comment LispValue = _comment(0)

func (_comment) Car() LispValue {
    return nil
}

func (_comment) Cdr() LispValue {
    return nil
}

func (_comment) IsNil() bool {
    return true
}

func (_comment) Repr() string {
    return ""
}

func IsComment(lv LispValue) bool {
    _, ok := lv.(_comment)
    return ok
}
