package executor

import (
	"github.com/corpix/schedulers/work"
)

type Executor interface {
	Execute(work.Work)
}
