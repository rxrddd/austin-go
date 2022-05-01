package content_model

import (
	"austin-go/app/austin-common/model"
	"austin-go/app/austin-common/types"
	"austin-go/common/zutils/dd"
	"testing"
)

func TestNewSmsContentModel(t *testing.T) {
	var messageTemplate = model.MessageTemplate{
		ID:         1,
		MsgContent: `{"content":"恭喜你:xxx:{$content}","url":"","title":""}`,
	}
	var messageParam = types.MessageParam{
		Receiver: "13788888888",
		Variables: map[string]string{
			"url":     "cnblogs.com/rainbow-tan/p/15628059.html",
			"content": "6666164180",
		},
		Extra: nil,
	}

	content := NewSmsContentModel().BuilderContent(messageTemplate, messageParam)
	dd.Print(content)
}
