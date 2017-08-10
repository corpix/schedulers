package schedulers

import (
	"fmt"
)

type ErrUnknownSchedulerType struct {
	t string
}

func (e *ErrUnknownSchedulerType) Error() string {
	return fmt.Sprintf(
		"Unknown scheduler type '%s'",
		e.t,
	)
}

func NewErrUnknownSchedulerType(t string) error {
	return &ErrUnknownSchedulerType{t}
}

//
