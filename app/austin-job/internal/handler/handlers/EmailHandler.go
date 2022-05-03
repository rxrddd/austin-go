package handlers

import (
	"austin-go/app/austin-common/dto/content_model"
	"austin-go/app/austin-common/types"
	"fmt"
)

type emailHandler struct {
}

func NewEmailHandler() IHandler {
	return emailHandler{}
}
func (h emailHandler) DoHandler(taskInfo types.TaskInfo) (err error) {
	fmt.Println(taskInfo)
	var content content_model.EmailContentModel
	getContentModel(taskInfo.ContentModel, &content)

	return nil
}
