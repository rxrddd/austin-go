package script

import (
	"austin-go/app/austin-common/types"
	"context"
)

type SmsScript interface {
	Send(ctx context.Context, taskInfo types.TaskInfo) (err error)
}

const TENCENT = "tencent" //腾讯云
const ALIYUN = "aliyun"   //阿里云
