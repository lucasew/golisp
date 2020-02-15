package vm_default

import (
    "github.com/lucasew/golisp/datatypes"
    common "github.com/lucasew/golisp/vm"
    "os"
    "reflect"
)

var env = map[string]datatypes.LispValue{}

// A function that depends of the environment to run
type LoadFunction func(common.LispVM) func(datatypes.LispValue) (datatypes.LispValue, error)

func NewDefaultEnv(parent *common.LispEnv) *common.LispEnv {
    env := map[string]interface{} {}
    // boolean constants
    env["nil"] = datatypes.Nil
    env["t"] = datatypes.T
    // system constants
    env["*TARGET*"] = []string{"golang", "default"}
    env["*HOSTNAME*"], _ = os.Hostname()
    // basic primitives
    env["car"] = func(v datatypes.LispValue) (datatypes.LispValue, error) {
        return v.Car(), nil
    }
    env["cdr"] = func(v datatypes.LispValue) (datatypes.LispValue, error) {
        return v.Cdr(), nil
    }
    // is-* functions
    env["is-number"] = func(v datatypes.LispValue) (datatypes.LispValue, error) {
        _, is := v.(datatypes.LispNumber)
        return boolToLispValue(is), nil
    }
    env["is-string"] = func(v datatypes.LispValue) (datatypes.LispValue, error) {
        _, is := v.(datatypes.LispString)
        return boolToLispValue(is), nil
    }
    env["is-symbol"] = func(v datatypes.LispValue) (datatypes.LispValue, error) {
        _, is := v.(datatypes.Symbol)
        return boolToLispValue(is), nil
    }
    env["is-function"] = func(v datatypes.LispValue) (datatypes.LispValue, error) {
        _, is := v.(datatypes.LispFunction)
        return boolToLispValue(is), nil
    }
    env["is-function-native"] = func(v datatypes.LispValue) (datatypes.LispValue, error) {
        f, is := v.(datatypes.LispFunction)
        if is {
            return boolToLispValue(f.IsFunctionNative()), nil
        }
        return datatypes.Nil, nil
    }
    env["is-atom"] = func(v datatypes.LispValue) (datatypes.LispValue, error) {
        _, is := v.(datatypes.Atom)
        return boolToLispValue(is), nil
    }
    env["is-cons"] = func(v datatypes.LispValue) (datatypes.LispValue, error) {
        _, is := v.(datatypes.Cons)
        return boolToLispValue(is), nil
    }
    // get length of something
    env["len"] = func(v datatypes.LispValue) (datatypes.LispValue, error) {
        if v.IsNil() {
            return datatypes.NewIntFromInt64(0), nil
        }
        t := reflect.TypeOf(v).Kind()
        if t == reflect.Array || t == reflect.Chan || t == reflect.Map || t == reflect.Slice || t == reflect.String {
            return datatypes.NewIntFromInt64(int64(reflect.ValueOf(v).Len())), nil
        }
        return datatypes.NewIntFromInt64(1), nil
    }
    env["not"] = func(v datatypes.LispValue) (datatypes.LispValue, error) {
        return boolToLispValue(v.IsNil()), nil // nil is false, then !true
    }
    // Returns the first non nil value
    //TODO: Test
    env["or"] = func(v datatypes.LispValue) (datatypes.LispValue, error) {
        if v.IsNil() {
            return datatypes.Nil, nil
        }
        begin:
        if !v.Car().IsNil() {
            return v.Car(), nil
        }
        v = v.Cdr()
        goto begin
    }
    env["and"] = func(v datatypes.LispValue) (datatypes.LispValue, error) {
        if v.IsNil() {
            return datatypes.Nil, nil
        }
        begin:
        if v.Car().IsNil() {
            return boolToLispValue(false), nil
        }
        if v.Cdr().IsNil() {
            return v.Car(), nil
        }
        v = v.Cdr()
        goto begin
    }
    e := common.NewLispEnv(parent)
    for k, v := range env {
        val, err := datatypes.NewLispValue(v)
        if err != nil {
            panic(err)
        }
        e.SetGlobal(k, val)
    }
    return e
}

func boolToLispValue(b bool) datatypes.LispValue {
    if b {
        return datatypes.T
    } else {
        return datatypes.Nil
    }
}
