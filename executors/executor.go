package executors

import (
	"strings"

	"github.com/corpix/schedulers/executors/executor"
	"github.com/corpix/schedulers/executors/executor/inplace"
	"github.com/corpix/schedulers/executors/executor/pool"
)

const (
	InplaceExecutorType = "inplace"
	PoolExecutorType    = "pool"
)

func NewFromConfig(c Config) (executor.Executor, error) {
	switch strings.ToLower(c.Type) {
	case InplaceExecutorType:
		return inplace.NewFromConfig(c.Inplace)
	case PoolExecutorType:
		return pool.NewFromConfig(c.Pool)
	default:
		return nil, NewErrUnknownExecutorType(c.Type)
	}
}
