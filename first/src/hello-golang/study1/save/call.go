package routes

import (
	"git.ot24.net/go/engine/errors"
	"model/boss"
	"model/user"
	"web/app/model/proto"
)

func init() {
	RegisterHandler("/call/transfer/set", SetCallTransfer)
	RegisterHandler("/call/transfer/get", GetCallTransfer)
}

var open string

// 实现设置呼转协议
func SetCallTransfer(req *proto.Request) *proto.Resp {
	fmt.Println(" 0 is close,1 is open")
	open := req.GetParam("open")
	var status int

	if open == "1" {
		status = 1
	} else {
		status = 2
	}

	uid := req.GetUid()
	u, err := user.UserByUid(uid)
	if err != nil {
		log.Warn(errors.As(err))
		return proto.NewResp(proto.ERR_UNKNOW, req.Ctx.Lang)
	}

	boss.SubcribeDtplan(u.MobileNum.Num, "21", status) //假定转移业务号为21
	balance, err := boss.balance(u.MobileNum.Num)
	if err != nil {
		log.Warn(errors.As(err))
		return proto.NewResp(proto.ERR_UNKNOW, req.Ctx.Lang)
	}
	return proto.NewResp(proto.SUCCESS, req.Ctx.Lang).
		PutData("balance", balance)
}

// 实现查询呼转协议
func GetCallTransfer(req *proto.Request) *proto.Resp {
	uid := req.GetUid()
	u, err := user.UserByUid(uid)

	return proto.NewResp(proto.ERR_UNKNOW, req.Ctx.Lang)
}
