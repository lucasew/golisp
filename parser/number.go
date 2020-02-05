package parser

import (
    "github.com/lucasew/golisp/datatypes"
    "github.com/lucasew/golisp/lex"
    "errors"
    "fmt"
    "strings"
)

func ParseNumber(ctx *lex.Context) (datatypes.LispValue, error) {
    b, ok := ctx.GetByte()
    if !ok {
        return datatypes.Nil, errors.New("eof when parsing number")
    }
    if !b.IsByteNumber() {
        return datatypes.Nil, errors.New("invalid entry point for number")
    }
    e := false // For scientific notation like 10^2 that is like 1E2
    dot := false // For the dot of floats
    slash := false // For the / in rational numbers
    begin := ctx.Index()
    for {
        ctx.Increment()
        b, ok := ctx.GetByte()
        if !ok {
            // return datatypes.Nil, errors.New("eof when parsing number body")
        }
        if b.IsByteE() {
            if e {
                return datatypes.Nil, errors.New("cant use the 'e' token more than one time to represent a number")
            }
            e = true
            continue
        }
        if b.IsDot() {
            if dot {
                return datatypes.Nil, errors.New("cant use the '.' token more than one time to represent a number")
            }
            dot = true
            continue
        }
        if b.IsSlash() {
            if slash {
                return datatypes.Nil, errors.New("cant use the '/' token more than one time to represent a number")
            }
            slash = true
            continue
        }
        if b.IsByteNumber() {
            continue
        }
        if b.IsByteUnderline() { // TODO: Ignore it when using the function to parse
            continue
        }
        s := ctx.Slice(begin, ctx.Index())
        s = strings.ReplaceAll(s, "_", "")
        reti, ok := datatypes.NewIntFromString(s)
        if ok {
            return reti, nil
        }
        retf, ok := datatypes.NewFloatFromString(s)
        if ok {
            return retf, nil
        }

        retr, ok := datatypes.NewRationalFromString(s)
        if ok {
            return retr, nil
        }
        return datatypes.Nil, fmt.Errorf("cant parse %s as number", s)
    }
}
