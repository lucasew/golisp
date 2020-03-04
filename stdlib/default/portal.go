package stdlib

import (
	"github.com/lucasew/golisp/data"
	"github.com/lucasew/golisp/data/types"
	"github.com/lucasew/golisp/data/types/number"
	"github.com/lucasew/golisp/utils/enforce"
)

func init() {
	register("is-portal", IsPortal)
	register("new-portal", NewPortal)
	register("portal-send", PortalSend)
	register("portal-send-unblocking", PortalSendUnblocking)
	register("portal-recv", PortalRecv)
	register("portal-recv-unblocking", PortalRecvUnblocking)
}

func NewPortal(v ...data.LispValue) (data.LispValue, error) {
	n := 0
	if len(v) != 0 {
		err := enforce.Validate(enforce.Length(v, 1), enforce.Integer(v, 1))
		if err != nil {
			return types.Nil, err
		}
		num, _ := v[0].(number.LispInt).Int64()
		n = int(num)
	}
	return types.NewPortal(n), nil
}

func IsPortal(v ...data.LispValue) (data.LispValue, error) {
	_, ok := v[0].(data.LispPortal)
	if ok {
		return types.T, nil
	} else {
		return types.Nil, nil
	}
}

func PortalSend(v ...data.LispValue) (data.LispValue, error) {
	err := enforce.Validate(enforce.Length(v, 2), enforce.Portal(v, 1))
	if err != nil {
		return types.Nil, err
	}
	p := v[0].(data.LispPortal)
	val := v[1]
	return p.Send(val), nil
}

func PortalSendUnblocking(v ...data.LispValue) (data.LispValue, error) {
	err := enforce.Validate(enforce.Length(v, 2), enforce.Portal(v, 1))
	if err != nil {
		return types.Nil, err
	}
	p := v[0].(data.LispPortal)
	val := v[1]
	return p.SendUnblocking(val), nil
}

func PortalRecv(v ...data.LispValue) (data.LispValue, error) {
	err := enforce.Validate(enforce.Length(v, 1), enforce.Portal(v, 1))
	if err != nil {
		return types.Nil, err
	}
	p := v[0].(data.LispPortal)
	return p.Recv(), nil
}

func PortalRecvUnblocking(v ...data.LispValue) (data.LispValue, error) {
	err := enforce.Validate(enforce.Length(v, 1), enforce.Portal(v, 1))
	if err != nil {
		return types.Nil, err
	}
	p := v[0].(data.LispPortal)
	return p.RecvUnblocking(), nil
}
