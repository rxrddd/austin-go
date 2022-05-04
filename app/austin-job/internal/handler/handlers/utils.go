package handlers

import (
	"github.com/zeromicro/go-zero/core/jsonx"
)

func getContentModel(contentModel interface{}, v interface{}) {
	marshal, _ := jsonx.Marshal(contentModel)
	_ = jsonx.Unmarshal(marshal, &v)
}

