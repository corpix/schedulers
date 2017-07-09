package task

import (
	"github.com/corpix/scheduler/schedule"
	"github.com/corpix/scheduler/work"
)

type Task struct {
	schedule.Schedule
	Fn work.Work
}

func New(s schedule.Schedule, fn work.Work) *Task {
	return &Task{s, fn}
}
