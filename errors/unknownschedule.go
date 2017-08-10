package errors

import (
	"fmt"
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
