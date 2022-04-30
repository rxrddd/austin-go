package interfaces

import "context"

type Process interface {
	Process(ctx context.Context, data interface{}) error
}
