package handlers

import (
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
	return nil
}
