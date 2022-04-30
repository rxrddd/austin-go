package action

import "austin-go/common/xerr"

var (
	sendErr         = xerr.NewErrMsg("发送消息错误")
	clientParamsErr = xerr.NewErrMsg("客户端参数错误")
)
