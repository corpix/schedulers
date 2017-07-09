package scheduler

import (
	"github.com/corpix/scheduler/executor"
	"github.com/corpix/scheduler/periodical"
	"github.com/corpix/scheduler/task"
)

type Scheduler interface {
	Schedule(*task.Task) error
	Unschedule(*task.Task)
	Close()
}

func NewFromConfig(e executor.Executor, config interface{}) (Scheduler, error) {
	switch v := config.(type) {
	case periodical.Config:
		return periodical.New(e, v)
	default:
		return nil, NewErrUnknownConfigType(config)
	}
}
