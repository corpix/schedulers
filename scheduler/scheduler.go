package scheduler

import (
	"github.com/corpix/schedulers/task"
)

type Scheduler interface {
	Schedule(*task.Task) error
	Unschedule(*task.Task)
	Close()
}
