package stdlib

import (
    "github.com/lucasew/golisp/data"
    "github.com/lucasew/golisp/data/types"
    "fmt"
)

func init() {
    register("is-portal", IsPortal)
    register("new-portal", NewPortal)
    register("portal-send", PortalSend)
    register("portal-send-unblocking", PortalSendUnblocking)
    register("portal-recv", PortalRecv)
    register("portal-recv-unblocking", PortalRecvUnblocking)
}

func NewPortal(v data.LispCons) (data.LispValue, error) {
    n, ok := v.Car().(types.LispInt)
    if v.Car().IsNil() {
        n = types.NewIntFromInt64(0)
    } else {
        if !ok {
            return types.Nil, fmt.Errorf("expected integer got %T", v.Car())
        }
    }
    num, _ := n.Int64()
    return types.NewPortal(int(num)), nil
}

func IsPortal(v data.LispCons) (data.LispValue, error) {
    _, ok := v.Car().(data.LispPortal)
    if ok {
        return types.T, nil
    } else {
        return types.Nil, nil
    }
}

func PortalSend(v data.LispCons) (data.LispValue, error) {
    p, ok := v.Car().(data.LispPortal)
    if !ok {
        return types.Nil, fmt.Errorf("first argument expected portal got %T", v.Car())
    }
    val := v.Cdr().Car()
    return p.Send(val), nil
}

func PortalSendUnblocking(v data.LispCons) (data.LispValue, error) {
    p, ok := v.Car().(data.LispPortal)
    if !ok {
        return types.Nil, fmt.Errorf("first argument expected portal got %T", v.Car())
    }
    val := v.Cdr().Car()
    return p.SendUnblocking(val), nil
}

func PortalRecv(v data.LispCons) (data.LispValue, error) {
    p, ok := v.Car().(data.LispPortal)
    if !ok {
        return types.Nil, fmt.Errorf("first argument expected portal got %T", v.Car())
    }
    return p.Recv(), nil
}

func PortalRecvUnblocking(v data.LispCons) (data.LispValue, error) {
    p, ok := v.Car().(data.LispPortal)
    if !ok {
        return types.Nil, fmt.Errorf("first argument expected portal got %T", v.Car())
    }
    return p.RecvUnblocking(), nil
}
