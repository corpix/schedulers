package executor

import (
	"github.com/corpix/scheduler/work"
)

type Executor interface {
	Execute(work.Work)
}
