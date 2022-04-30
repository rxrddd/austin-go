package content_model

import "austin-go/app/austin-common/interfaces"

var contentMap = map[int]interfaces.BuilderContent{
	30: NewSmsContentModel(),
	//50: NewSmsContentModel(),
}

//消息发送渠道：10.IM 20.Push 30.短信 40.Email 50.公众号 60.小程序 70.企业微信
func GetBuilderContentBySendChannel(sendChannel int) interfaces.BuilderContent {
	return contentMap[sendChannel]
}
