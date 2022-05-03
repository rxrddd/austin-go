package content_model

import (
	"austin-go/app/austin-common/model"
	"austin-go/app/austin-common/types"
	"austin-go/common/zutils/dd"
	"fmt"
	"testing"
)

func TestNewSmsContentModel(t *testing.T) {
	var messageTemplate = model.MessageTemplate{
		ID:         1,
		MsgContent: `{"content":"恭喜你:xxx:{$content}","url":"","title":""}`,
	}
	var messageParam = types.MessageParam{
		Receiver: "13788888888",
		Variables: map[string]interface{}{
			"url":     "cnblogs.com/rainbow-tan/p/15628059.html",
			"content": "6666164180",
			"map": map[string]string{
				"appid": "xxx",
			},
		},
	}

	content := NewSmsContentModel().BuilderContent(messageTemplate, messageParam)
	dd.Print(content)
}

func TestXX(t *testing.T) {
	var a = make(map[string]string)
	a["name"] = "张三"
	var b string
	fmt.Println(a["name"])
	fmt.Println(a["李四"])
	b = a["李四"]
	fmt.Println(b)
}
