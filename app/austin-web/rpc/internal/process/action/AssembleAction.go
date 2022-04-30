package action

import (
	"austin-go/app/austin-common/content_model"
	"austin-go/app/austin-common/task_util"
	"austin-go/app/austin-common/types"
	"austin-go/common/zutils/transform"
	"austin-go/repo"
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"strings"
)

type AssembleAction struct {
}

func NewAssembleAction() *AssembleAction {
	return &AssembleAction{}
}

func (p AssembleAction) Process(ctx context.Context, data interface{}) error {
	logx.Info(data)
	sendTaskModel, ok := data.(*types.SendTaskModel)
	if !ok {
		return errors.Wrapf(sendErr, "AssembleAction 类型错误 err:%v", data)
	}
	messageParamList := sendTaskModel.MessageParamList

	messageTemplate, err := repo.NewMessageTemplateRepo().One(ctx, sendTaskModel.MessageTemplateId)
	if err != nil {
		return errors.Wrapf(sendErr, "查询模板异常 err:%v 模板id:%d", err, sendTaskModel.MessageTemplateId)
	}
	contentModel := content_model.GetBuilderContentBySendChannel(messageTemplate.SendChannel)

	var newTaskList []types.TaskInfo
	for _, param := range messageParamList {

		curTask := types.TaskInfo{
			MessageTemplateId: messageTemplate.ID,
			BusinessId:        task_util.GenerateBusinessId(messageTemplate.ID, messageTemplate.TemplateType),
			Receiver:          transform.ArrayStringUniq(strings.Split(param.Receiver, ",")),
			IdType:            messageTemplate.IDType,
			SendChannel:       messageTemplate.SendChannel,
			TemplateType:      messageTemplate.TemplateType,
			MsgType:           messageTemplate.MsgType,
			ShieldType:        messageTemplate.ShieldType,
			ContentModel:      contentModel.BuilderContent(*messageTemplate, param),
			SendAccount:       messageTemplate.SendAccount,
		}

		newTaskList = append(newTaskList, curTask)
	}
	sendTaskModel.TaskInfo = newTaskList
	return nil
}
