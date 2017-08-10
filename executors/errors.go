package executors

import (
	"fmt"
)

type ErrUnknownExecutorType struct {
	t string
}

func (e *ErrUnknownExecutorType) Error() string {
	return fmt.Sprintf(
		"Unknown executor type '%s'",
		e.t,
	)
}

func NewErrUnknownExecutorType(t string) error {
	return &ErrUnknownExecutorType{t}
}
