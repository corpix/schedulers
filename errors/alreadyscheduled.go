package errors

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"

	"github.com/corpix/schedulers/task"
)

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
