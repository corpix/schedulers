package executor

import (
	"fmt"
)

type ErrUnknownConfigType struct {
	config interface{}
}

func (e *ErrUnknownConfigType) Error() string {
	return fmt.Sprintf(
		"Unknown config type '%T'",
		e.config,
	)
}
func NewErrUnknownConfigType(config interface{}) error {
	return &ErrUnknownConfigType{config}
}
