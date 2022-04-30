package interfaces

import (
	"austin-go/app/austin-common/model"
	"austin-go/app/austin-common/types"
)

type BuilderContent interface {
	BuilderContent(messageTemplate model.MessageTemplate, messageParam types.MessageParam) interface{}
}
