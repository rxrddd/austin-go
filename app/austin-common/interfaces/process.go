package interfaces

import (
	"austin-go/app/austin-common/types"
	"context"
)

type Process interface {
	Process(ctx context.Context, sendTaskModel *types.SendTaskModel) error
}
