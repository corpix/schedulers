package task

import (
	"github.com/corpix/schedulers/schedule"
	"github.com/corpix/schedulers/work"
)

type Task struct {
	schedule.Schedule
	Fn work.Work
}

func New(s schedule.Schedule, fn work.Work) *Task {
	return &Task{s, fn}
}
