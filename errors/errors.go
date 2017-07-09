package errors

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"

	"github.com/corpix/scheduler/task"
)

type ErrUnknownSchedule struct {
	want interface{}
	got  interface{}
}

func (e *ErrUnknownSchedule) Error() string {
	return fmt.Sprintf(
		"Unknown schedule type, want '%T' got '%T'",
		e.want,
		e.got,
	)
}
func NewErrUnknownSchedule(want, got interface{}) error {
	return &ErrUnknownSchedule{want, got}
}

//

type ErrAlreadyScheduled struct {
	task *task.Task
}

func (e *ErrAlreadyScheduled) Error() string {
	return fmt.Sprintf(
		"Task '%s' already scheduled",
		spew.Sdump(e.task),
	)
}
func NewErrAlreadyScheduled(t *task.Task) error {
	return &ErrAlreadyScheduled{t}
}

//
