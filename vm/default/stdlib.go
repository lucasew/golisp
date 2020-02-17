package vm_default

import (
    "github.com/lucasew/golisp/data"
    "github.com/lucasew/golisp/data/convert"
    "github.com/lucasew/golisp/data/types"
    common "github.com/lucasew/golisp/vm"
    envpkg "github.com/lucasew/golisp/vm/components/env"
    "os"
    "reflect"
    "errors"
    "strings"
)

var scope = map[string]data.LispValue{}

// A function that depends of the environment to run
type LoadFunction func(common.LispVM) func(data.LispValue) (data.LispValue, error)

func NewDefaultEnv(parent *envpkg.LispEnv) *envpkg.LispEnv {
    e := map[string]interface{} {}
    // boolean constants
    e["nil"] = data.Nil
    e["t"] = data.T
    // system constants
    e["*TARGET*"] = []string{"golang", "default"}
    e["*HOSTNAME*"], _ = os.Hostname()
    // char constants
    e["char-bell"] = "\a"
    e["char-cr"] = "\r"
    e["char-nl"] = "\n"
    // basic primitives
    e["car"] = func(v data.LispValue) (data.LispValue, error) {
        return v.Car(), nil
    }
    e["cdr"] = func(v data.LispValue) (data.LispValue, error) {
        return v.Cdr(), nil
    }
    // is-* functions
    e["is-number"] = func(v data.LispValue) (data.LispValue, error) {
        _, is := v.(data.LispNumber)
        return boolToLispValue(is), nil
    }
    e["is-string"] = func(v data.LispValue) (data.LispValue, error) {
        _, is := v.(data.LispString)
        return boolToLispValue(is), nil
    }
    e["is-symbol"] = func(v data.LispValue) (data.LispValue, error) {
        _, is := v.(types.Symbol)
        return boolToLispValue(is), nil
    }
    e["is-function"] = func(v data.LispValue) (data.LispValue, error) {
        _, is := v.(data.LispFunction)
        return boolToLispValue(is), nil
    }
    e["is-function-native"] = func(v data.LispValue) (data.LispValue, error) {
        f, is := v.(data.LispFunction)
        if is {
            return boolToLispValue(f.IsFunctionNative()), nil
        }
        return data.Nil, nil
    }
    e["is-atom"] = func(v data.LispValue) (data.LispValue, error) {
        _, is := v.(types.Atom)
        return boolToLispValue(is), nil
    }
    e["is-cons"] = func(v data.LispValue) (data.LispValue, error) {
        _, is := v.(types.Cons)
        return boolToLispValue(is), nil
    }
    // get length of something
    e["len"] = func(v data.LispValue) (data.LispValue, error) {
        if v.IsNil() {
            return types.NewIntFromInt64(0), nil
        }
        t := reflect.TypeOf(v).Kind()
        if t == reflect.Array || t == reflect.Chan || t == reflect.Map || t == reflect.Slice || t == reflect.String {
            return types.NewIntFromInt64(int64(reflect.ValueOf(v).Len())), nil
        }
        return types.NewIntFromInt64(1), nil
    }
    e["not"] = func(v data.LispValue) (data.LispValue, error) {
        return boolToLispValue(v.IsNil()), nil // nil is false, then !true
    }
    // Returns the first non nil value
    //TODO: Test
    e["or"] = func(v data.LispValue) (data.LispValue, error) {
        if v.IsNil() {
            return data.Nil, nil
        }
        begin:
        if !v.Car().IsNil() {
            return v.Car(), nil
        }
        v = v.Cdr()
        goto begin
    }
    e["and"] = func(v data.LispValue) (data.LispValue, error) {
        if v.IsNil() {
            return data.Nil, nil
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
    e["print"] = func(v data.LispValue) (data.LispValue, error) {
        str := []string{}
        for val := v.Car(); !v.IsNil(); v = v.Cdr() {
            println(v.Repr())
            s, ok := val.(data.LispString)
            if !ok {
                return data.Nil, errors.New("I can only print strings, to print other types convert it first")
            }
            str = append(str, s.ToString())
        }
        s := strings.Join(str, "")
        print(s)
        return types.NewConventionalString(s), nil
    }
    e["println"] = func(v data.LispValue) (data.LispValue, error) {
        str := []string{}
        for val := v.Car(); !v.IsNil(); v = v.Cdr() {
            s, ok := val.(types.ConventionalString)
            if !ok {
                return data.Nil, errors.New("I can only print strings, to print other types convert it first")
            }
            str = append(str, string(s))
        }
        s := strings.Join(str, "")
        println(s)
        return types.NewConventionalString(s), nil
    }
    r := envpkg.NewLispEnv(parent)
    for k, v := range e {
        val, err := convert.NewLispValue(v)
        if err != nil {
            panic(err)
        }
        r.SetGlobal(k, val)
    }
    return r
}

func boolToLispValue(b bool) data.LispValue {
    if b {
        return data.T
    } else {
        return data.Nil
    }
}
