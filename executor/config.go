package executor

import (
	"github.com/corpix/scheduler/executor/inplace"
	"github.com/corpix/scheduler/executor/pool"
)

func NewFromConfig(c interface{}) (Executor, error) {
	switch v := c.(type) {
	case pool.Config:
		return pool.NewFromConfig(v)
	case inplace.Config:
		return inplace.NewFromConfig(v)
	default:
		return nil, NewErrUnknownConfigType(c)
	}
}
