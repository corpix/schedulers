package executors

import (
	"github.com/corpix/schedulers/executors/executor/inplace"
	"github.com/corpix/schedulers/executors/executor/pool"
)

type Config struct {
	Type    string
	Inplace inplace.Config
	Pool    pool.Config
}
