package utils

import "github.com/zeromicro/go-zero/core/jsonx"

func GetContentModel(contentModel interface{}, v interface{}) {
	marshal, _ := jsonx.Marshal(contentModel)
	_ = jsonx.Unmarshal(marshal, &v)
}
