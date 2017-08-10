package schedules

import (
	"fmt"
)

type ErrUnknownScheduleType struct {
	t string
}

func (e *ErrUnknownScheduleType) Error() string {
	return fmt.Sprintf(
		"Unknown schedule type '%s'",
		e.t,
	)
}
func NewErrUnknownScheduleType(t string) error {
	return &ErrUnknownScheduleType{t}
}
