package process

import (
	"austin-go/app/austin-web/rpc/internal/process/interfaces"
	"context"
)

type BusinessProcess struct {
	process []interfaces.Process
}

func NewBusinessProcess() *BusinessProcess {
	return &BusinessProcess{
		process: make([]interfaces.Process, 0),
	}
}
func (p *BusinessProcess) Process(ctx context.Context, data interface{}) error {
	for _, pr := range p.process {
		err := pr.Process(ctx, data)
		if err != nil {
			return err
		}
	}
	return nil
}
func (p *BusinessProcess) AddProcess(pr ...interfaces.Process) error {
	if len(pr) > 0 {
		p.process = append(p.process, pr...)
	}
	return nil
}
